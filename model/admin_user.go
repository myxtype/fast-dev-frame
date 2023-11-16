package model

import (
	"frame/pkg/sql/types"
	"gorm.io/gorm"
	"time"
)

// 管理员
type AdminUser struct {
	gorm.Model
	Username    string `gorm:"uniqueIndex;size:255"`
	Password    types.Password
	LastLoginAt *time.Time
}
