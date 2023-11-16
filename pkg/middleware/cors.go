package middleware

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"time"
)

func Cors() gin.HandlerFunc {
	return cors.New(cors.Config{
		AllowOrigins:     []string{"*"},
		AllowMethods:     []string{"GET", "POST", "PUT", "PATCH", "DELETE", "OPTIONS"},
		AllowHeaders:     []string{"Accept", "Origin", "XRequestedWith", "Content-Type", "LastModified", "X-Access-Token", "X-Lang"},
		AllowCredentials: true,
		MaxAge:           365 * 24 * time.Hour,
	})
}
