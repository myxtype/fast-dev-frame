package admin

import (
	"frame/model"
	"frame/pkg/sql/sqltypes"
	"time"
)

type PageRequest struct {
	Limit int               `json:"pageSize" form:"pageSize"`
	Page  int               `json:"current" form:"current"`
	Sort  map[string]string `json:"sort" form:"sort"`
}

type ListResult struct {
	Total int64       `json:"total"`
	Data  interface{} `json:"data"`
}

func NewListResult(total int64, data interface{}) *ListResult {
	return &ListResult{
		Total: total,
		Data:  data,
	}
}

// 管理员
type AdminUserInfoVo struct {
	ID        uint      `json:"id"`
	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"`
	Username  string    `json:"username"`
	RoleId    uint      `json:"roleId"`   // 角色
	Name      string    `json:"name"`     // 昵称
	Avatar    string    `json:"avatar"`   // 头像
	Disabled  bool      `json:"disabled"` // 是否禁用
}

func NewAdminUserInfoVo(v *model.AdminUser) *AdminUserInfoVo {
	return &AdminUserInfoVo{
		ID:        v.ID,
		CreatedAt: v.CreatedAt,
		UpdatedAt: v.UpdatedAt,
		Username:  v.Username,
		RoleId:    v.RoleId,
		Name:      v.Name,
		Avatar:    v.Avatar,
		Disabled:  v.Disabled,
	}
}

// 管理员
type AdminUserVo struct {
	ID           uint      `json:"id"`
	CreatedAt    time.Time `json:"createdAt"`
	UpdatedAt    time.Time `json:"updatedAt"`
	Username     string    `json:"username"`
	RoleId       uint      `json:"roleId"`       // 角色
	Name         string    `json:"name"`         // 昵称
	Avatar       string    `json:"avatar"`       // 头像
	Disabled     bool      `json:"disabled"`     // 是否禁用
	LoginVersion int64     `json:"loginVersion"` // 登录版本
}

func NewAdminUserVo(v *model.AdminUser) *AdminUserVo {
	return &AdminUserVo{
		ID:        v.ID,
		CreatedAt: v.CreatedAt,
		UpdatedAt: v.UpdatedAt,
		Username:  v.Username,
		RoleId:    v.RoleId,
		Name:      v.Name,
		Avatar:    v.Avatar,
		Disabled:  v.Disabled,
	}
}

// 管理员角色
type AdminRoleVo struct {
	ID          uint                 `json:"id"`
	CreatedAt   time.Time            `json:"createdAt"`
	UpdatedAt   time.Time            `json:"updatedAt"`
	Name        string               `json:"name"`        // 角色名称
	Permissions sqltypes.StringArray `json:"permissions"` // 权限列表
	Disabled    bool                 `json:"disabled"`    // 是否禁用
}

func NewAdminRoleVo(v *model.AdminRole) *AdminRoleVo {
	return &AdminRoleVo{
		ID:          v.ID,
		CreatedAt:   v.CreatedAt,
		UpdatedAt:   v.UpdatedAt,
		Name:        v.Name,
		Permissions: v.Permissions,
		Disabled:    v.Disabled,
	}
}
