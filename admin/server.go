package admin

import (
	"frame/pkg/grace"
	"frame/pkg/middleware"
	"github.com/gin-gonic/gin"
	"io/ioutil"
	"net/http"
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

	// 建议定一个特殊的前缀
	x := r.Group("/x")
	{
		x.GET("/ping", func(ctx *gin.Context) {
			ctx.String(http.StatusOK, "ok")
		})
	}

	// 优雅的重启
	grace.HttpRun(server.addr, r, 10*time.Minute)
}
