package config

import (
	"github.com/joho/godotenv"
	"log"
	"os"
)

var (
	Host     = os.Getenv("DB_HOST")     // Адрес БД
	Port     = os.Getenv("DB_PORT")     // Порт
	User     = os.Getenv("DB_USER")     // Имя пользователя
	Password = os.Getenv("DB_PASSWORD") // Пароль
	DbName   = os.Getenv("DB_NAME")     // Имя базы
)

func LoadEnv() {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatal("Error loading .env file")
	}
}
