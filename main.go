package main

import (
	"github.com/gin-gonic/gin"
	"go-template/app/controller"
	"go-template/app/middleware"
	"os"
)

import "net/http"

func main() {
	router := gin.Default()

	// アプリケーションのポート番号を環境変数から取得します。
	port := os.Getenv("APP_PORT")
	if port == "" {
		port = "8080" // デフォルトのポート番号
	}
	router.Use(middleware.RecordUaAndTime)
	router.Use(middleware.Cors())
	router.Use(func(c *gin.Context) {
		if c.Request.Method != "OPTIONS" {
			middleware.ValidateKey()(c)
		}
	})

	router.Use(middleware.Cors())
	router.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "hello world",
		})
	})

	novelEngine := router.Group("/novel")
	{
		v1 := novelEngine.Group("/v1")
		{
			v1.GET("/id/:id", controller.NovelById)
			v1.POST("/add", controller.NovelAdd)
			v1.GET("/list", controller.NovelList)
			v1.PUT("/update", controller.NovelUpdate)
			v1.DELETE("/delete", controller.NovelDelete)
		}
	}
	router.Run(":" + port)
}
