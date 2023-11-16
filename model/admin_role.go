package model

import (
	"frame/pkg/sql/sqltypes"
	"gorm.io/gorm"
)

// 管理员角色
type AdminRole struct {
	gorm.Model
	Name        string               // 角色名称
	Permissions sqltypes.StringArray // 权限列表
	Disabled    bool                 // 是否禁用
}
