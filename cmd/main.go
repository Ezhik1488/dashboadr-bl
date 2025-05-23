package main

import (
	"dashboard-bl/config"
	"dashboard-bl/internal/src"
	"fmt"
	"github.com/gofiber/fiber/v2"
	"github.com/sirupsen/logrus"
	"os"
	"time"
)

func main() {

	fmt.Println(time.Now())

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
	src.SetupRoutes(app, logger)

	// Запуск сервера
	err = app.Listen(":3000")
	if err != nil {
		return
	}
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
