package model

import (
	"gorm.io/gorm"
)

// AdminLog 管理员操作日志
type AdminLog struct {
	gorm.Model
	AdminId uint   `gorm:"index"` // 管理员ID
	Notes   string // 备注
	Ip      string // IP
}
