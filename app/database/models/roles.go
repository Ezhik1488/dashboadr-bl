package models

type Role struct {
	ID       int    `gorm:"primaryKey;autoIncrement" json:"id"`
	RoleName string `gorm:"size:50;not null" json:"role_name"`
}
