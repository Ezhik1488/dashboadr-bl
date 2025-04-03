package models

import "time"

type ClientDocument struct {
	ID           int       `json:"id"`
	ClientID     int       `gorm:"foreignKey:ClientID;not null" json:"client_id"`
	Pdf          string    `gorm:"type:text" json:"pdf"`
	DocumentType string    `gorm:"size:255;not null" json:"document_type"`
	CreatedAt    time.Time `gorm:"autoCreateTime" json:"created_at"`
	UpdatedAt    time.Time `gorm:"autoUpdateTime" json:"updated_at"`
}
