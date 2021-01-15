package rest

import (
	"frame/pkg/ecode"
	"frame/pkg/request"
	"frame/service"
	"github.com/gin-gonic/gin"
	"github.com/spf13/cast"
)

// 获取用户数据
func GetUserByUserId(ctx *gin.Context) {
	app := request.New(ctx)

	id := cast.ToInt64(ctx.Query("id"))
	user, err := service.UserService.GetUserById(id)
	if err != nil {
		app.Response(err)
		return
	}

	app.Response(nil, NewUserVo(user))
}

type UserRegisterRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

// 用户注册
func UserRegister(ctx *gin.Context) {
	app := request.New(ctx)

	var req UserRegisterRequest
	if err := ctx.BindJSON(&req); err != nil {
		app.Response(err)
		return
	}

	if len(req.Username) == 0 || len(req.Password) == 0 {
		app.Response(ecode.ErrRequest)
		return
	}

	err := service.UserService.Register(req.Username, req.Password)
	if err != nil {
		app.Response(err)
		return
	}

	app.Response(nil)
}
