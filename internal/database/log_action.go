package database

import (
	"dashboard-bl/internal/database/models"
	"gorm.io/gorm"
)

func LogAction(db *gorm.DB, userID int, details, actionType string) error {
	var action = models.UserAction{
		UserID:            userID,
		ActionType:        actionType,
		ActionDescription: details,
	}

	return db.Create(&action).Error
}
