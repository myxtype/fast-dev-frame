package rest

import (
	"frame/pkg/ecode"
	"frame/pkg/request"
	"frame/service"
	"github.com/gin-gonic/gin"
)

type UserController struct{}

type GetUserRequest struct {
	ID uint `json:"id" form:"id" binding:"required"`
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
	Username string `json:"username" form:"username" binding:"required"`
	Password string `json:"password" form:"username" binding:"required"`
}

func (*UserController) UserRegister(ctx *gin.Context) {
	app := request.New(ctx)

	var req UserRegisterRequest
	if err := ctx.ShouldBind(&req); err != nil {
		app.Error(err)
		return
	}

	err := service.UserService.Register(ctx, req.Username, req.Password)
	if err != nil {
		app.Error(err)
		return
	}

	app.Success(nil)
}
