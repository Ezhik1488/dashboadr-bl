package src

import (
	"dashboard-bl/graph"
	"dashboard-bl/internal/database"
	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/handler/transport"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/gofiber/fiber/v2"
	"github.com/sirupsen/logrus"
	"github.com/valyala/fasthttp/fasthttpadaptor"
	"gorm.io/gorm"
)

func SetupRoutes(app *fiber.App, logger *logrus.Logger) {
	db := database.GetConnection()

	initGraphQL(app, db, logger)

}

func initGraphQL(app *fiber.App, db *gorm.DB, logger *logrus.Logger) {
	// clientRepo := clients.NewClientRepo(db, logger)
	// clientService := clients.NewClientServices(db, logger, clientRepo)

	resolver := &graph.Resolver{}
	schema := handler.New(graph.NewExecutableSchema(graph.Config{Resolvers: resolver}))

	// Настройка поддерживаемых транспортов
	schema.AddTransport(transport.Options{})
	schema.AddTransport(transport.GET{})
	schema.AddTransport(transport.POST{})
	schema.AddTransport(transport.MultipartForm{})

	// Обработчик GraphQL
	app.All("/graphql", func(c *fiber.Ctx) error {
		fasthttpadaptor.NewFastHTTPHandlerFunc(schema.ServeHTTP)(c.Context())
		return nil
	})

	app.Get("/", func(c *fiber.Ctx) error {
		fasthttpadaptor.NewFastHTTPHandlerFunc(
			playground.Handler("GraphQL Playground", "/graphql").ServeHTTP,
		)(c.Context())
		return nil
	})
}
