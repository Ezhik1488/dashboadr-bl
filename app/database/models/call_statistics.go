package models

import "time"

type CallStatistic struct {
	ID          int       `gorm:"primary_key;AUTO_INCREMENT" json:"id"`
	ClientId    int       `gorm:"foreignKey:ClientId;not null" json:"client_id"`
	MinutesUsed float64   `gorm:"type:decimal(10,2);not null" json:"minutes_used"`
	TotalCount  float64   `gorm:"type:decimal(10,2);not null" json:"total_count"`
	Month       int       `gorm:"type:integer;not null" json:"month"`
	Year        int       `gorm:"type:integer;not null" json:"year"`
	CreatedAt   time.Time `gorm:"autoCreateTime" json:"created_at"`
}
