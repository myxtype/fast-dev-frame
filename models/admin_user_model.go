package models

import "time"

// 管理员
type AdminUser struct {
	ID          int64 `gorm:"primaryKey;autoIncrement"`
	CreatedAt   time.Time
	UpdatedAt   time.Time
	Username    string `gorm:"uniqueIndex;size:255"`
	Password    string
	LastLoginAt *time.Time
}
