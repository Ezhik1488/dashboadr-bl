package models

import "time"

type PhoneNumber struct {
	ID           int       `gorm:"primary_key;AUTO_INCREMENT" json:"id"`
	ClientID     int       `gorm:"foreignKey:ClientID" json:"client_id"`
	Number       string    `gorm:"size:15" json:"number"`
	AssignedDate time.Time `gorm:"type:TIMESTAMP" json:"assigned_date"`
	MonthlyPrice float64   `gorm:"type:decimal(10,2);not null" json:"monthly_price"`
	CreatedAt    time.Time `gorm:"autoCreateTime" json:"created_at"`
	UpdatedAt    time.Time `gorm:"autoUpdateTime" json:"updated_at"`
}
