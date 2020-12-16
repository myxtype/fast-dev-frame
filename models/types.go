package models

import (
	"database/sql/driver"
	"frame/pkg/sql/format"
)

// 用户身份
type UserIdentity string

const (
	UserIdentityNormal = UserIdentity("normal") // 普通
	UserIdentityAgent  = UserIdentity("agent")  // 代理商
)

// 用户信息
type UserInfo struct {
	Sex   int8   `json:"sex"`   // 性别：0保密，1男，2女
	Email string `json:"email"` // 邮件
}

func (t *UserInfo) Scan(src interface{}) error {
	return format.Scan(t, src)
}

func (t UserInfo) Value() (driver.Value, error) {
	return format.Value(t)
}
