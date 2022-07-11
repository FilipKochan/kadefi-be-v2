package controllers

import (
	"net/http"

	"github.com/FilipKochan/kadefi-be-v2/models"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func GetPools(ctx *gin.Context, db *gorm.DB) {
	pools := []models.Pool{}

	err := db.Preload("Platform").Preload("Token").Find(&pools).Error
	if err != nil {
		ctx.String(http.StatusInternalServerError, err.Error())
		return
	}
	ctx.IndentedJSON(http.StatusOK, pools)
}

func GetPool(ctx *gin.Context, db *gorm.DB) {
	pool := models.Pool{}

	if err := ctx.BindUri(&pool); err != nil {
		ctx.String(http.StatusBadRequest, err.Error())
		return
	}

	if err := db.Preload("Platform").Preload("Token").First(&pool).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			ctx.String(http.StatusNotFound, err.Error())
			return
		}

		ctx.String(http.StatusInternalServerError, err.Error())
		return
	}

	ctx.IndentedJSON(http.StatusOK, pool)
}

func GetPoolByPlatformToken(ctx *gin.Context, db *gorm.DB) {
	pool := models.Pool{}

	if err := ctx.BindUri(&pool); err != nil {
		ctx.String(http.StatusBadRequest, err.Error())
		return
	}

	if err := db.Preload("Platform").Preload("Token").Where("platform_id = ?", pool.PlatformID).Where("token_id = ?", pool.TokenID).First(&pool).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			ctx.String(http.StatusNotFound, err.Error())
			return
		}

		ctx.String(http.StatusInternalServerError, err.Error())
		return
	}

	ctx.IndentedJSON(http.StatusOK, pool)
}
