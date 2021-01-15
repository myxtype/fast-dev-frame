package model

import (
	"frame/pkg/sql/types"
	"time"
)

// 用户表
type User struct {
	ID        int64 `gorm:"primaryKey;autoIncrement"`
	CreatedAt time.Time
	UpdatedAt time.Time
	Username  string         `gorm:"uniqueIndex;size:255"` // 用户名
	Password  types.Password // 登录密码hash
}
