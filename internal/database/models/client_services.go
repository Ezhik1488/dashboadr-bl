package models

import "time"

type ClientService struct {
	ID        int       `gorm:"primary_key;AUTO_INCREMENT" json:"id"`
	ClientID  int       `gorm:"foreignKey:ClientID" json:"client_id"`
	ServiceID int       `gorm:"foreignKey:ServiceID" json:"service_id"`
	StartDate time.Time `gorm:"default:CURRENT_TIMESTAMP" json:"start_date"`
	EndDate   time.Time `gorm:"type:TIMESTAMP" json:"end_date"`

	// Связи
	Service Service `gorm:"foreignKey:ServiceID" json:"service,omitempty"`
}
