package rest

import (
	"frame/pkg/grace"
	"frame/pkg/middleware"
	"github.com/gin-gonic/gin"
	"io/ioutil"
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
	gin.DefaultWriter = ioutil.Discard

	r := gin.Default()
	r.Use(middleware.SetCROSOptions)

	v1 := r.Group("/v1")
	{
		v1.GET("/user", GetUserByUserId)
		v1.POST("/user", UserRegister)
	}

	// 优雅的重启
	grace.HttpRun(server.addr, r, 10*time.Minute)
}
