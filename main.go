package main

import (
	"github.com/gin-gonic/gin"
	"go-template/app/middleware"
	"os"
)

import "net/http"

func main() {
	// gin.Default() は Logger と Recovery ミドルウェアを付けた *Engine を返します。
	router := gin.Default()

	// CORS ミドルウェアを使うことで、クロスオリジンリソース共有 (CORS) を許可します。
	router.Use(middleware.Cors())

	// アプリケーションのポート番号を環境変数から取得します。
	port := os.Getenv("APP_PORT")

	router.Use(middleware.Cors())
	router.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "hello world",
		})
	})
	router.Run(":" + port)
}
