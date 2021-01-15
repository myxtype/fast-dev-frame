package app

import "github.com/gin-gonic/gin"

type App struct {
	c *gin.Context
}

func New(c *gin.Context) *App {
	return &App{c: c}
}
