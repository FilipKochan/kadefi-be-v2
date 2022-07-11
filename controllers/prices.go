package controllers

import (
	"net/http"

	"github.com/FilipKochan/kadefi-be-v2/models"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func GetPrices(ctx *gin.Context, db *gorm.DB) {
	prices := []models.PriceRecord{}

	if err := db.Limit(2).Preload("Pool").Preload("Pool.Platform").Preload("Pool.Token").Find(&prices).Error; err != nil {
		ctx.String(http.StatusInternalServerError, err.Error())
		return
	}
	ctx.IndentedJSON(http.StatusOK, prices)
}

func PostPrice(ctx *gin.Context, db *gorm.DB) {
	price := models.PriceRecord{}

	if err := ctx.BindJSON(&price); err != nil {
		ctx.String(http.StatusBadRequest, err.Error())
		return
	}

	price.ID = 0

	newest := models.PriceRecord{}
	if err := db.Where("pool_id", price.PoolID).Order("updated_at DESC").First(&newest).Error; err != nil && err != gorm.ErrRecordNotFound {
		ctx.String(http.StatusInternalServerError, err.Error())
		return
	}

	if newest.PriceInKda == price.PriceInKda {
		if err := db.Save(newest).Error; err != nil {
			ctx.String(http.StatusInternalServerError, err.Error())
			return
		}
		db.Preload("Pool").Preload("Pool.Platform").Preload("Pool.Token").First(&newest)
		ctx.IndentedJSON(http.StatusOK, newest)
		return
	}

	if err := db.Create(&price).Error; err != nil {
		ctx.String(http.StatusInternalServerError, err.Error())
		return
	}

	db.Preload("Pool").Preload("Pool.Platform").Preload("Pool.Token").First(&price)
	ctx.IndentedJSON(http.StatusCreated, &price)
}
