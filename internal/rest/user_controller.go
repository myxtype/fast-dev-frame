package rest

import (
	"frame/pkg/ecode"
	"frame/pkg/request"
	"frame/service"
	"github.com/gin-gonic/gin"
)

type UserController struct{}

type GetUserRequest struct {
	ID uint `json:"id" form:"id"`
}

func (*UserController) GetUser(ctx *gin.Context) {
	app := request.New(ctx)

	var req GetUserRequest
	if err := ctx.ShouldBind(&req); err != nil {
		app.Error(err)
		return
	}

	user, err := service.UserService.GetUserByID(ctx, req.ID)
	if err != nil {
		app.Error(err)
		return
	}

	if user == nil {
		app.Error(ecode.ErrNotFind)
		return
	}

	app.Success(NewUserVo(user))
}

type UserRegisterRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

func (*UserController) UserRegister(ctx *gin.Context) {
	app := request.New(ctx)

	var req UserRegisterRequest
	if err := ctx.BindJSON(&req); err != nil {
		app.Error(err)
		return
	}

	if len(req.Username) == 0 || len(req.Password) == 0 {
		app.Error(ecode.ErrRequest)
		return
	}

	err := service.UserService.Register(ctx, req.Username, req.Password)
	if err != nil {
		app.Error(err)
		return
	}

	app.Success(nil)
}
