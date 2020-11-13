package middleware

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

// 跨域
func SetCROSOptions(c *gin.Context) {
	c.Header("Access-Control-Allow-Origin", "*")
	c.Header("Access-Control-Allow-Methods", "GET,POST,PUT,PATCH,DELETE,OPTIONS")
	c.Header("Access-Control-Allow-Headers", "Accept,Origin,XRequestedWith,Content-Type,LastModified,X-Access-Token,X-Lang,X-Api-Key")
	c.Header("Access-Control-Max-Age", "1800")
	c.Header("Allow", "HEAD,GET,POST,PUT,PATCH,DELETE,OPTIONS")

	if c.Request.Method == http.MethodOptions {
		c.AbortWithStatus(http.StatusOK)
	}
}
