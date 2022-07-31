package main

import (
	"fmt"
	"log"
	"os"

	"github.com/FilipKochan/kadefi-be-v2/controllers"
	"github.com/FilipKochan/kadefi-be-v2/models"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

const defaultPort string = "8081"

func main() {
	router := setupRouter()
	router.Run(getPort())
}

func setupRouter() *gin.Engine {
	godotenv.Load(".env")

	// gin.SetMode(gin.ReleaseMode)
	router := gin.Default()
	router.SetTrustedProxies([]string{os.Getenv("APP_URL")})
	db, err := dbConnect()

	if err != nil {
		log.Fatal(err)
		panic(err.Error())
	}

	models.Migrate(db)

	pingController := controllers.PingController{}
	indexController := controllers.IndexController{}
	platformController := controllers.PlatformsController{}
	tokensController := controllers.TokensController{}
	poolsController := controllers.PoolsController{}
	kdaUsdRatesController := controllers.KdaUsdRatesController{}
	pricesController := controllers.PricesController{}

	router.GET("/ping", pingController.Get)
	api := router.Group("/")
	{
		api.GET("", func(ctx *gin.Context) {
			indexController.Get(ctx)
		})
		platforms := api.Group("/platforms")
		{
			platforms.GET("", func(ctx *gin.Context) {
				platformController.Get(ctx, db)
			})
			platforms.GET(":id", func(ctx *gin.Context) {
				platformController.GetOne(ctx, db)
			})
		}
		tokens := api.Group("/tokens")
		{
			tokens.GET("", func(ctx *gin.Context) {
				tokensController.Get(ctx, db)
			})
			tokens.GET(":id", func(ctx *gin.Context) {
				tokensController.GetOne(ctx, db)
			})
		}
		pools := api.Group("/pools")
		{
			pools.GET("", func(ctx *gin.Context) {
				poolsController.Get(ctx, db)
			})
			pools.GET("/platforms/:platformId/tokens/:tokenId", func(ctx *gin.Context) {
				poolsController.GetByPlatformToken(ctx, db)
			})
			pools.GET(":id", func(ctx *gin.Context) {
				poolsController.GetOne(ctx, db)
			})
		}
		kdaUsdRates := api.Group("/kda-usd-rates")
		{
			kdaUsdRates.GET("", func(ctx *gin.Context) {
				kdaUsdRatesController.Get(ctx, db)
			})
			kdaUsdRates.POST("", func(ctx *gin.Context) {
				kdaUsdRatesController.Post(ctx, db)
			})
		}
		prices := api.Group("/prices")
		{
			prices.GET("", func(ctx *gin.Context) {
				pricesController.Get(ctx, db)
			})
			prices.POST("", func(ctx *gin.Context) {
				pricesController.Post(ctx, db)
			})
		}
	}

	return router
}

func dbConnect() (*gorm.DB, error) {
	dsn := fmt.Sprintf("%v:%v@tcp(%v:3306)/%v?charset=utf8mb4&parseTime=True&loc=Local",
		os.Getenv("DB_USERNAME"),
		os.Getenv("DB_PASSWORD"),
		os.Getenv("DB_HOST"),
		os.Getenv("DB_NAME"))

	return gorm.Open(mysql.Open(dsn), &gorm.Config{})
}

func getPort() string {
	port_ := os.Getenv("PORT")
	if port_ == "" {
		return ":" + defaultPort
	}
	return ":" + port_
}
