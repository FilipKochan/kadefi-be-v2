package models

type Pool struct {
	BaseModel
	Volume     float64  `json:"volume"`
	PlatformID uint     `gorm:"not null" json:"platformId" uri:"platformId"`
	Platform   Platform `json:"platform"`
	TokenID    uint     `gorm:"not null" json:"tokenId" uri:"tokenId"`
	Token      Token    `json:"token"`
}
