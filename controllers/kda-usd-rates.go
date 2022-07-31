package controllers

import (
	"net/http"

	"github.com/FilipKochan/kadefi-be-v2/models"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type KdaUsdRatesController struct{}

func (_ KdaUsdRatesController) Get(ctx *gin.Context, db *gorm.DB) {
	rates := []models.KdaToUsdRate{}

	if err := db.Find(&rates).Error; err != nil {
		ctx.String(http.StatusInternalServerError, err.Error())
		return
	}

	ctx.IndentedJSON(http.StatusOK, rates)
}

func (_ KdaUsdRatesController) Post(ctx *gin.Context, db *gorm.DB) {
	rate := models.KdaToUsdRate{}

	if err := ctx.BindJSON(&rate); err != nil {
		ctx.String(http.StatusBadRequest, err.Error())
		return
	}

	rate.ID = 0

	newest := models.KdaToUsdRate{}
	if err := db.Order("updated_at DESC").First(&newest).Error; err != nil && err != gorm.ErrRecordNotFound {
		ctx.String(http.StatusInternalServerError, err.Error())
		return
	}

	if newest.KdaToUsd == rate.KdaToUsd {
		if err := db.Save(&newest).Error; err != nil {
			ctx.String(http.StatusInternalServerError, err.Error())
			return
		}

		ctx.IndentedJSON(http.StatusOK, &newest)
		return
	}

	if err := db.Create(&rate).Error; err != nil {
		ctx.String(http.StatusBadRequest, err.Error())
		return
	}

	ctx.IndentedJSON(http.StatusCreated, &rate)
}
