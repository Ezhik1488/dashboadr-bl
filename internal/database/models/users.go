package models

import "time"

type User struct {
	ID           int       `gorm:"primaryKey;autoIncrement" json:"id"`
	Login        string    `gorm:"size:30;not null;unique" json:"login"`
	Email        string    `gorm:"size:100;not null;unique" json:"email"`
	PasswordHash string    `gorm:"size:255;not null" json:"password"`
	Roles        []Role    `gorm:"many2many:user_roles;" json:"roles"`
	LastLogin    time.Time `json:"last_login"`
	CreatedAt    time.Time `gorm:"autoCreateTime" json:"created_at"`
	UpdatedAt    time.Time `gorm:"autoUpdateTime" json:"updated_at"`
}
