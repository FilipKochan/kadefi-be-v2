package models

type KdaToUsdRate struct {
	BaseModel
	KdaToUsd float64 `gorm:"not null" json:"kdaToUsd"`
}
