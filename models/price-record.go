package models

type PriceRecord struct {
	BaseModel
	PriceInKda     float64 `gorm:"not null" json:"priceInKda"`
	TokenLiquidity float64 `gorm:"not null" json:"tokenLiquidity"`
	PoolID         uint    `gorm:"not null" json:"poolId"`
	Pool           Pool    `json:"pool"`
}
