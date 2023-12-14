package rest

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

		user, err := service.UserService.CheckToken(ctx, token)
		if err != nil {
			app.Error(err)
			return
		}

		app.SetUser(user)
		ctx.Next()
	}
}
