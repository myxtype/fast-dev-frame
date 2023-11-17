package rest

import (
	"frame/pkg/grace"
	"frame/pkg/middleware"
	"github.com/gin-contrib/cache"
	"github.com/gin-contrib/cache/persistence"
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

	// 轻量级缓存
	store := persistence.NewInMemoryStore(time.Second)

	r := gin.Default()
	r.Use(middleware.Cors())

	v1 := r.Group("/v1")
	{
		userGroup := v1.Group("/user")
		{
			user := &UserController{}
			userGroup.GET("/user", cache.CachePage(store, 1*time.Minute, user.GetUser))
			userGroup.POST("/user", user.UserRegister)
		}
	}

	// 优雅地重启
	grace.HttpRun(server.addr, r, 10*time.Minute)
}
