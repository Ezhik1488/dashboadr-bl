package main

import (
	"dashboard-bl/config"
	"dashboard-bl/graph"
	"dashboard-bl/internal/src"
	"fmt"
	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/gofiber/fiber/v2"
	"github.com/sirupsen/logrus"
	"github.com/valyala/fasthttp/fasthttpadaptor"
	"net/http"
	"os"
)

func main() {
	app := fiber.New(fiber.Config{
		Prefork: false,
	})

	// Загрузка переменных окружения
	config.LoadEnv()

	// Инициализация логгера приложения
	logger, err := setupLogger(config.LogPath)
	if err != nil {
		logrus.Fatal(err)
	}

	// Установка маршрутов
	src.SetupRoutes(logger)

	// Инициализация GraphQL
	initGraphQL(app)

	// Запуск сервера
	err = app.Listen(":3000")
	if err != nil {
		return
	}
}

func initGraphQL(app *fiber.App) {
	graphqlHandler := handler.New(graph.NewExecutableSchema(graph.Config{Resolvers: &graph.Resolver{}}))
	app.Post("/graphql", func(c *fiber.Ctx) error {
		fasthttpadaptor.NewFastHTTPHandler(http.HandlerFunc(graphqlHandler.ServeHTTP))(c.Context())
		return nil
	})

	// GraphQL Playground
	app.Get("/", func(c *fiber.Ctx) error {
		fasthttpadaptor.NewFastHTTPHandler(playground.Handler("GraphQL Playground", "/graphql"))(c.Context())
		return nil
	})
}

func setupLogger(logPath string) (*logrus.Logger, error) {
	logger := logrus.New()

	// Открываем файл для записи логов
	file, err := os.OpenFile(logPath, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0644)
	if err != nil {
		return nil, fmt.Errorf("failed to open log file: %w", err)
	}

	logger.SetOutput(file)
	lvl, err := logrus.ParseLevel(config.LogLevel)
	if err != nil {
		lvl = logrus.InfoLevel
	}
	logger.SetLevel(lvl)
	// logger.SetFormatter(&logrus.JSONFormatter{})

	return logger, nil
}
