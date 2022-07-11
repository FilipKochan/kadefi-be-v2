package controllers

import (
	"net/http"

	"github.com/FilipKochan/kadefi-be-v2/models"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func GetPlatforms(ctx *gin.Context, db *gorm.DB) {
	platforms := []models.Platform{}
	err := db.Find(&platforms).Error
	if err != nil {
		ctx.String(http.StatusInternalServerError, err.Error())
		return
	}

	ctx.IndentedJSON(http.StatusOK, platforms)
}

func GetPlatform(ctx *gin.Context, db *gorm.DB) {
	platform := models.Platform{}

	if err := ctx.ShouldBindUri(&platform); err != nil {
		ctx.String(http.StatusBadRequest, err.Error())
		return
	}

	if err := db.First(&platform).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			ctx.String(http.StatusNotFound, err.Error())
			return
		}

		ctx.String(http.StatusInternalServerError, err.Error())
		return
	}

	ctx.IndentedJSON(http.StatusOK, platform)
}
