package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetPing(c *gin.Context) {
	c.String(http.StatusOK, "pong")
}

func GetIndex(c *gin.Context) {
	c.String(http.StatusOK, "Backend for <a href='http://kadefi.app'>kadefi.app</a>")
}
