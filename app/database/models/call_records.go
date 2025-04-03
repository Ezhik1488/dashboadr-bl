package models

import "time"

type CallRecord struct {
	ID       int       `gorm:"primary_key;AUTO_INCREMENT" json:"id"`
	ClientID int       `gorm:"foreignKey:ClientID;not null" json:"client_id"`
	CallType string    `gorm:"size:10;not null" json:"call_type"`
	ANumber  string    `gorm:"size:20;not null" json:"a_number"`
	BNumber  string    `gorm:"size:20;not null" json:"b_number"`
	Duration int       `gorm:"default:0" json:"duration"`
	CallDate time.Time `gorm:"type:TIMESTAMP;not null" json:"call_date"`
}
