package database

import (
	"dashboard-bl/config"
	"dashboard-bl/internal/database/models"
	"fmt"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"log"
)

func GetConnection() *gorm.DB {
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable",
		config.Host, config.User, config.Password, config.DbName, config.Port)
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}
	// Миграция схем
	dbErr := db.AutoMigrate(
		&models.User{},
		&models.Client{},
		&models.ClientDetail{},
		&models.CallStatistic{},
		&models.ClientDocument{},
		&models.ClientService{},
		&models.PhoneNumber{},
		&models.Service{},
		&models.Role{},
		&models.UserAction{},
		&models.CallRecord{},
	)
	if dbErr != nil {
		log.Fatal("Не удалось выполнить миграцию таблицы")
	}
	return db
}
