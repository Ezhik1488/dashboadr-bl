package models

import "time"

type UserAction struct {
	ID                int       `gorm:"primaryKey;autoIncrement" json:"id"`
	UserID            int       `gorm:"foreignKey:UserID" json:"user_id"`
	ActionType        string    `gorm:"size:50;not null" json:"action_type"`
	ActionDescription string    `gorm:"type:text;not null" json:"action_description"`
	CreatedAt         time.Time `gorm:"default:CURRENT_TIMESTAMP" json:"action_time"`
}
