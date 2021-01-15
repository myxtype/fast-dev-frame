package rest

import (
	"frame/pkg/app"
	"frame/pkg/ecode"
	"frame/service"
	"github.com/gin-gonic/gin"
	"github.com/spf13/cast"
)

// 获取用户数据
func GetUserByUserId(ctx *gin.Context) {
	appG := app.New(ctx)

	id := cast.ToInt64(ctx.Query("id"))
	user, err := service.UserService.GetUserById(id)
	if err != nil {
		appG.Response(err)
		return
	}

	appG.Response(nil, NewUserVo(user))
}

type UserRegisterRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

// 用户注册
func UserRegister(ctx *gin.Context) {
	appG := app.New(ctx)

	var req UserRegisterRequest
	if err := ctx.BindJSON(&req); err != nil {
		appG.Response(err)
		return
	}

	if len(req.Username) == 0 || len(req.Password) == 0 {
		appG.Response(ecode.ErrRequest)
		return
	}

	err := service.UserService.Register(req.Username, req.Password)
	if err != nil {
		appG.Response(err)
		return
	}

	appG.Response(nil)
}
