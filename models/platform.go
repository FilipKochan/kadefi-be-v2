package models

type Platform struct {
	BaseModel
	Name string `gorm:"not null" json:"name"`
}
