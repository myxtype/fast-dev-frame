package admin

import (
	"frame/pkg/ecode"
	"frame/pkg/request"
	"frame/service"
	"github.com/gin-gonic/gin"
	"time"
)

type AuthController struct{}

type AuthLoginRequest struct {
	Username string `json:"username" form:"username"`
	Password string `json:"password" form:"password"`
}

func (c *AuthController) Login(ctx *gin.Context) {
	app := request.New(ctx)

	var req AuthLoginRequest
	if err := ctx.Bind(&req); err != nil {
		app.Error(err)
		return
	}

	if req.Username == "" || req.Password == "" {
		app.Error(ecode.ErrRequest)
		return
	}

	count, err := service.AdminService.GetAdminUserCount(ctx)
	if err != nil {
		app.Error(err)
		return
	}

	if count == 0 {
		if err := service.AdminService.InitAdmin(ctx); err != nil {
			app.Error(err)
			return
		}
	}

	user, token, err := service.AdminService.Login(ctx, req.Username, req.Password)
	if err != nil {
		app.Error(err)
		return
	}

	service.AdminService.AddLog(ctx, user.ID, "登录成功", ctx.ClientIP())

	app.Success(token)
}

func (c *AuthController) OutLogin(ctx *gin.Context) {
	time.Sleep(time.Second)
	request.New(ctx).Success(nil)
}
