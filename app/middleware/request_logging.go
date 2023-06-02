package middleware

import (
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"log"
	"os"
	"time"
)

func RecordUaAndTime(c *gin.Context) {
	logFile, err := os.OpenFile("logs/request.log", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		log.Fatal(err)
	}
	encoder := zapcore.NewJSONEncoder(zap.NewProductionEncoderConfig())
	core := zapcore.NewTee(
		zapcore.NewCore(encoder, zapcore.AddSync(logFile), zapcore.InfoLevel),
		zapcore.NewCore(encoder, zapcore.AddSync(os.Stdout), zapcore.InfoLevel),
	)
	logger := zap.New(core)

	oldTime := time.Now()
	ua := c.GetHeader("User-Agent")
	c.Next()
	logger.Info("incoming request",
		zap.String("path", c.Request.URL.Path),
		zap.String("method", c.Request.Method),
		zap.String("ip", c.ClientIP()),
		zap.String("Ua", ua),
		zap.Any("headers", c.Request.Header),
		zap.Int("status", c.Writer.Status()),
		zap.Duration("elapsed", time.Now().Sub(oldTime)),
	)
}
