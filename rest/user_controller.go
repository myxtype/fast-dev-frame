package rest

import (
	"frame/pkg/ecode"
	"frame/pkg/response"
	"frame/service"
	"github.com/gin-gonic/gin"
	"github.com/spf13/cast"
	"net/http"
)

// 获取用户数据
func GetUserByUserId(ctx *gin.Context) {
	id := cast.ToInt64(ctx.Query("id"))

	user, err := service.UserService.GetUserById(id)
	if err != nil {
		ctx.JSON(http.StatusOK, response.JSON(err))
		return
	}

	ctx.JSON(http.StatusOK, response.JSON(nil, NewUserVo(user)))
}

type UserRegisterRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

// 用户注册
func UserRegister(ctx *gin.Context) {
	var req UserRegisterRequest
	if err := ctx.BindJSON(&req); err != nil {
		ctx.JSON(http.StatusOK, response.JSON(err))
		return
	}

	if len(req.Username) == 0 || len(req.Password) == 0 {
		ctx.JSON(http.StatusOK, response.JSON(ecode.ErrRequest))
		return
	}

	err := service.UserService.Register(req.Username, req.Password)
	if err != nil {
		ctx.JSON(http.StatusOK, response.JSON(err))
		return
	}

	ctx.JSON(http.StatusOK, response.JSON(nil))
}
