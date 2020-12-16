package models

import (
	"time"
)

// 用户表
type User struct {
	ID        int64 `gorm:"primaryKey;autoIncrement"`
	CreatedAt time.Time
	UpdatedAt time.Time
	Username  string       `gorm:"uniqueIndex;size:255"` // 用户名
	Password  string       // 登录密码hash
	Identity  UserIdentity `gorm:"size:255"`  // 用户身份
	Info      UserInfo     `gorm:"type:json"` // 用户信息
}

// 管理员
type AdminUser struct {
	ID          int64 `gorm:"primaryKey;autoIncrement"`
	CreatedAt   time.Time
	UpdatedAt   time.Time
	Username    string `gorm:"uniqueIndex;size:255"`
	Password    string
	LastLoginAt *time.Time
}
