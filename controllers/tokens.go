package controllers

import (
	"net/http"

	"github.com/FilipKochan/kadefi-be-v2/models"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func GetTokens(ctx *gin.Context, db *gorm.DB) {
	tokens := []models.Token{}
	err := db.Find(&tokens).Error
	if err != nil {
		ctx.String(http.StatusInternalServerError, err.Error())
		return
	}

	ctx.IndentedJSON(http.StatusOK, &tokens)
}

func GetToken(ctx *gin.Context, db *gorm.DB) {
	token := models.Token{}

	if err := ctx.ShouldBindUri(&token); err != nil {
		ctx.String(http.StatusBadRequest, err.Error())
		return
	}

	if err := db.First(&token).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			ctx.String(http.StatusNotFound, err.Error())
			return
		}

		ctx.String(http.StatusInternalServerError, err.Error())
		return
	}

	ctx.IndentedJSON(http.StatusOK, &token)
}
