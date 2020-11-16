package models

import (
	"time"
)

// 用户表
type User struct {
	ID        int64 `gorm:"primaryKey;autoIncrement"`
	CreatedAt time.Time
	UpdatedAt time.Time
	Username  string `gorm:"uniqueIndex"`
	Password  string
}
