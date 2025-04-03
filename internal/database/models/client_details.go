package models

type ClientDetail struct {
	ID           int    `gorm:"primaryKey;autoIncrement" json:"id"`
	ClientID     int    `gorm:"foreignKey:ClientID" json:"client_id"`
	FullName     string `gorm:"size:100;default:'not specified'" json:"full_name"`
	Director     string `gorm:"size:100;default:'not specified'" json:"director"`
	Kpp          string `gorm:"size:30;default:'not specified'" json:"kpp"`
	Ogrn         string `gorm:"size:30;default:'not specified'" json:"ogrn"`
	LegalAddress string `gorm:"size:100;default:'not specified'" json:"legal_address"`
	PostAddress  string `gorm:"size:100;default:'not specified'" json:"post_address"`
	Bank         string `gorm:"size:100;default:'not specified'" json:"bank"`
	Bik          string `gorm:"size:20;default:'not specified'" json:"bik"`
	CorrAcc      string `gorm:"size:40;default:'not specified'" json:"corr_acc"`
	CalcAcc      string `gorm:"size:40;default:'not specified'" json:"calc_acc"`
}
