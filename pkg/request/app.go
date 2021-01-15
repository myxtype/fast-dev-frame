package request

import "github.com/gin-gonic/gin"

type AppRequest struct {
	c *gin.Context
}

func New(c *gin.Context) *AppRequest {
	return &AppRequest{c: c}
}
