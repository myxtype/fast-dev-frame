package rest

import (
	"frame/models"
	"frame/pkg/ecode"
	"time"
)

// 响应的数据结构
type ResponseVo struct {
	Message string      `json:"message"`
	Code    int         `json:"code"`
	Data    interface{} `json:"data,omitempty"`
}

// 数据响应
func NewResponseVo(err error, args ...interface{}) *ResponseVo {
	var data interface{}
	if len(args) > 0 {
		data = args[0]
	}

	ec := ecode.Cause(err)

	return &ResponseVo{
		Code:    ec.Code(),
		Message: ec.Message(),
		Data:    data,
	}
}

type UserVo struct {
	ID        int64     `json:"id"`
	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"`
	Username  string    `json:"username"`
}

func NewUserVo(v *models.User) *UserVo {
	return &UserVo{
		ID:        v.ID,
		CreatedAt: v.CreatedAt,
		UpdatedAt: v.UpdatedAt,
		Username:  v.Username,
	}
}
