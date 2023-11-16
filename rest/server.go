package rest

import (
	"frame/pkg/grace"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"io"
	"time"
)

type HttpServer struct {
	addr string
}

func NewHttpServer(addr string) *HttpServer {
	return &HttpServer{
		addr: addr,
	}
}

// 启动服务
func (server *HttpServer) Start() {
	gin.SetMode(gin.ReleaseMode)
	gin.DefaultWriter = io.Discard

	r := gin.Default()
	r.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"*"},
		AllowMethods:     []string{"GET", "POST", "PUT", "PATCH", "DELETE", "OPTIONS"},
		AllowHeaders:     []string{"Accept", "Origin", "XRequestedWith", "Content-Type", "LastModified", "X-Access-Token", "X-Lang"},
		AllowCredentials: true,
		MaxAge:           365 * 24 * time.Hour,
	}))

	v1 := r.Group("/v1")
	{
		v1.GET("/user", GetUserByUserId)
		v1.POST("/user", UserRegister)
	}

	// 优雅地重启
	grace.HttpRun(server.addr, r, 10*time.Minute)
}
