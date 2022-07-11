package models

type Token struct {
	BaseModel
	Name   string `gorm:"not null" json:"name"`
	Ticker string `gorm:"not null" json:"ticker"`
}
