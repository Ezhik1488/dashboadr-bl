package models

import "time"

type User struct {
	ID        int       `gorm:"primaryKey;autoIncrement" json:"id"`
	Login     string    `gorm:"size:30;not null;unique" json:"login"`
	Email     string    `gorm:"size:100;not null;unique" json:"email"`
	Password  string    `gorm:"size:255;not null" json:"password"`
	LastLogin time.Time `gorm:"type:timestamp" json:"last_login"`
	CreatedAt time.Time `gorm:"type:timestamp" json:"-"`
	UpdatedAt time.Time `gorm:"type:timestamp" json:"-"`
}
