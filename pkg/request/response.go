package request

import (
	"frame/pkg/ecode"
	"net/http"
)

// 统一响应的结构
type response struct {
	Message string      `json:"message"`
	Code    int         `json:"code"`
	Data    interface{} `json:"data,omitempty"`
}

// Response JSON响应
func (a *AppRequest) Response(err error, args ...interface{}) {
	var data interface{}
	if len(args) > 0 {
		data = args[0]
	}

	ec := ecode.Cause(err)
	a.c.JSON(http.StatusOK, &response{
		Code:    ec.Code(),
		Message: ec.Message(),
		Data:    data,
	})
}

// Success 成功响应
func (a *AppRequest) Success(args interface{}) {
	a.Response(nil, args)
}

// Error 错误响应
func (a *AppRequest) Error(err error) {
	a.c.Abort()
	a.Response(err)
}
