package ecode

var (
	ErrUserPassword = New(1000, "登录密码错误")
	ErrUserDisabled = New(1001, "账号已被禁用")
)
