package admin

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

	// 建议定一个特殊的前缀
	x := r.Group("/x")
	{
		authGroup := x.Group("/auth")
		{
			auth := &AuthController{}
			authGroup.POST("/login", auth.Login)
			authGroup.POST("/out-login", checkGhost(), auth.OutLogin)
		}

		pri := x.Group("/", checkToken())
		{
			adminGroup := pri.Group("/admin")
			{
				admin := &AdminController{}
				adminGroup.GET("/current", admin.Current)
				adminGroup.POST("/password", admin.UpdatePassword)

				adminGroup.GET("/roles", permit("Admin_Manage"), admin.QueryAdminRoles)
				adminGroup.GET("/list", permit("Admin_Manage"), admin.QueryAdminUsers)
				adminGroup.POST("/user", permit("Admin_Manage"), admin.SaveAdminUser)
				adminGroup.POST("/role", permit("Admin_Manage"), admin.SaveAdminRole)
			}
		}
	}

	// 优雅的重启
	grace.HttpRun(server.addr, r, 10*time.Minute)
}
