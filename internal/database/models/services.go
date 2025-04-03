package models

import "time"

type Service struct {
	ID          int       `gorm:"primary_key;AUTO_INCREMENT" json:"id"`
	Name        string    `gorm:"size:100;not null" json:"name"`
	Description string    `gorm:"default:'not specified'" json:"description"`
	MonthlyCost float64   `gorm:"type:decimal(10,2);not null" json:"monthly_cost"`
	CreatedAt   time.Time `gorm:"autoCreateTime" json:"created_at"`
	UpdatedAt   time.Time `gorm:"autoUpdateTime" json:"updated_at"`
}
