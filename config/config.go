package config

import (
	"github.com/joho/godotenv"
	_ "github.com/joho/godotenv/autoload"
	"log"
	"os"
)

var (
	Host     = os.Getenv("DB_HOST")     // Адрес БД
	Port     = os.Getenv("DB_PORT")     // Порт
	User     = os.Getenv("DB_USER")     // Имя пользователя
	Password = os.Getenv("DB_PASSWORD") // Пароль
	DbName   = os.Getenv("DB_NAME")     // Имя базы
	LogPath  = os.Getenv("LOG_PATH")    // Путь до файла логов
	LogLevel = os.Getenv("LOG_LEVEL")   // Уровень логирования
)

func LoadEnv() {
	err_ := godotenv.Load()

	// Проверка наличия необходимых переменных
	requiredVars := []string{"DB_HOST", "DB_PORT", "DB_USER", "DB_PASSWORD", "DB_NAME", "LOG_PATH"}
	for _, key := range requiredVars {
		if os.Getenv(key) == "" {
			log.Fatalf("Required environment variable %s is not set", key)
		}
	}

	if err_ != nil {
		log.Fatal("Error loading .env file", err_)
	}
}
