package models

import (
	"time"

	"gorm.io/gorm"
)

func Migrate(db *gorm.DB) {
	db.AutoMigrate(&KdaToUsdRate{})
	db.AutoMigrate(&Platform{})
	db.AutoMigrate(&Token{})
	db.AutoMigrate(&Pool{})
	db.AutoMigrate(&PriceRecord{})
}

type BaseModel struct {
	ID        uint           `gorm:"primarykey" uri:"id" json:"id"`
	CreatedAt time.Time      `json:"createdAt"`
	UpdatedAt time.Time      `json:"updatedAt"`
	DeletedAt gorm.DeletedAt `gorm:"index" json:"deletedAt"`
}
