package models

type UserRole struct {
	ID     int  `gorm:"primaryKey;autoIncrement" json:"id"`
	UserID User `gorm:"foreignKey:UserID;type:integer" json:"user_id"`
	RoleID Role `gorm:"foreignKey:RoleID;type:integer" json:"role_id"`
}
