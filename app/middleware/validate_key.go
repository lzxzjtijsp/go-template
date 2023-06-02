package middleware

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"strings"
)

func ValidateKey() gin.HandlerFunc {
	return func(c *gin.Context) {
		// リクエストヘッダーから Authorization ヘッダーの値を取得
		authHeader := c.GetHeader("Authorization")

		// Authorization ヘッダーが存在し、Bearer トークンが正しい形式であるか確認
		if authHeader != "" && strings.HasPrefix(authHeader, "Bearer ") {
			apiKey := strings.TrimPrefix(authHeader, "Bearer ")

			// 期待する API キーと一致するか確認
			if apiKey == "Your-Expected-Value" {
				c.Next()
				return
			}
		}

		// キーが見つからない場合や期待する値を持っていない場合の処理
		c.AbortWithStatus(http.StatusUnauthorized)
		return
	}
}
