package response

import "frame/pkg/ecode"

// 响应的数据结构
type ResponseJSON struct {
	Message string      `json:"message"`
	Code    int         `json:"code"`
	Data    interface{} `json:"data,omitempty"`
}

func JSON(err error, args ...interface{}) *ResponseJSON {
	var data interface{}
	if len(args) > 0 {
		data = args[0]
	}

	ec := ecode.Cause(err)
	return &ResponseJSON{
		Code:    ec.Code(),
		Message: ec.Message(),
		Data:    data,
	}
}
