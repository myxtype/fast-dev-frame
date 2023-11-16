package model

import (
	"frame/pkg/sql/types"
	"gorm.io/gorm"
)

// 用户表
type User struct {
	gorm.Model
	Username string         `gorm:"uniqueIndex;size:255"` // 用户名
	Password types.Password // 登录密码hash
}
