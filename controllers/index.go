package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type IndexController struct{}

type PingController struct{}

func (_ PingController) Get(c *gin.Context) {
	c.String(http.StatusOK, "pong")
}

func (_ IndexController) Get(c *gin.Context) {
	c.String(http.StatusOK, "Backend for <a href='http://kadefi.app'>kadefi.app</a>")
}
