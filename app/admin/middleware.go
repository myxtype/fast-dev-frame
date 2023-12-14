package admin

import (
	"frame/pkg/ecode"
	"frame/pkg/request"
	"frame/service"
	"github.com/gin-gonic/gin"
)

// 检查并获取登录用户
func checkToken() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		app := request.New(ctx)

		token := ctx.GetHeader("X-Access-Token")
		if token == "" {
			app.Error(ecode.ErrAuth)
			return
		}

		user, err := service.AdminService.CheckToken(ctx, token)
		if err != nil {
			app.Error(err)
			return
		}

		app.SetAdminUser(user)
		ctx.Next()
	}
}

// 检查游客登录信息
func checkGhost() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		app := request.New(ctx)

		token := ctx.GetHeader("X-Access-Token")
		if token != "" {
			user, err := service.AdminService.CheckToken(ctx, token)
			if err == nil {
				app.SetAdminUser(user)
			}
		}

		ctx.Next()
	}
}

// 检查权限
func permit(permit string) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		app := request.New(ctx)

		user := app.GetAdminUser()
		if user == nil || user.RoleId <= 0 {
			app.Error(ecode.ErrForbidden)
			return
		}

		pass, err := service.AdminService.CheckAdminRole(ctx, user, permit)
		if err != nil {
			app.Error(err)
			return
		}

		if !pass {
			app.Error(ecode.ErrNoPermission)
			return
		}

		ctx.Next()
	}
}
