package models

import "time"

type Client struct {
	ID          int       `gorm:"primary_key;AUTO_INCREMENT" json:"id"`
	Name        string    `gorm:"size:100;not null" json:"name"`
	Inn         string    `gorm:"size:20;not null" json:"inn"`
	Email       string    `gorm:"size:100;not null" json:"email"`
	IsKik       bool      `gorm:"default:false" json:"is_kik"`
	IsPaymaster bool      `gorm:"default:false" json:"is_paymaster"`
	CreatedAt   time.Time `gorm:"autoCreateTime" json:"created_at"`
	UpdatedAt   time.Time `gorm:"autoUpdateTime" json:"updated_at"`

	// Связи
	Details  ClientDetail `gorm:"foreignKey:ClientID" json:"details"`
	Services []Service    `gorm:"many2many:client_services" json:"services"`
}
