package src

import (
	"dashboard-bl/graph"
	"dashboard-bl/internal/database"
	"dashboard-bl/internal/domain/clients"
	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/gofiber/fiber/v2"
	"github.com/sirupsen/logrus"
	"github.com/valyala/fasthttp/fasthttpadaptor"
	"gorm.io/gorm"
	"net/http"
)

func SetupRoutes(app *fiber.App, logger *logrus.Logger) {
	db := database.GetConnection()

	initGraphQL(app, db, logger)

}

func initGraphQL(app *fiber.App, db *gorm.DB, logger *logrus.Logger) {
	clientRepo := clients.NewClientRepo(db, logger)
	clientService := clients.NewClientServices(db, logger, clientRepo)

	resolver := graph.NewResolver(&clientService)

	graphqlHandler := handler.New(graph.NewExecutableSchema(graph.Config{Resolvers: resolver}))
	app.Post("/graphql", func(c *fiber.Ctx) error {
		// Логирование входящего запроса
		logger.Infof("Incoming GraphQL request: %s %s", c.Method(), c.Path())

		// Обработка запроса
		fasthttpadaptor.NewFastHTTPHandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			graphqlHandler.ServeHTTP(w, r)
		})(c.Context())

		return nil
	})

	// GraphQL Playground
	app.Get("/", func(c *fiber.Ctx) error {

		fasthttpadaptor.NewFastHTTPHandler(playground.Handler("GraphQL Playground", "/graphql"))(c.Context())
		return nil
	})
}
