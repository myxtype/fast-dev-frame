package model

import (
	"frame/pkg/sql/sqltypes"
	"gorm.io/gorm"
)

// User 用户表
type User struct {
	gorm.Model
	Username string            `gorm:"uniqueIndex;size:255"` // 用户名
	Password sqltypes.Password // 登录密码hash
	Disabled bool              // 是否禁用
}
