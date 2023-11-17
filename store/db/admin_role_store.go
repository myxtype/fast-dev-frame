package db

import (
	"context"
	"frame/model"
	"gorm.io/gorm"
)

// GetAdminRoleById 查询管理员
func (s *Store) GetAdminRoleById(ctx context.Context, id uint) (*model.AdminRole, error) {
	var role model.AdminRole
	err := s.db.WithContext(ctx).First(&role, id).Error
	if err == gorm.ErrRecordNotFound {
		return nil, nil
	}
	return &role, err
}

// QueryAdminRoles 查询管理员角色权限
func (s *Store) QueryAdminRoles(ctx context.Context, page, limit int) ([]*model.AdminRole, int64, error) {
	db := s.db.WithContext(ctx)

	var data []*model.AdminRole
	var count int64
	err := db.Offset((page - 1) * limit).Limit(limit).Find(&data).Offset(-1).Limit(-1).Count(&count).Error

	return data, count, err
}

// SaveAdminRole 保存管理员角色权限
func (s *Store) SaveAdminRole(ctx context.Context, v *model.AdminRole) error {
	return s.db.WithContext(ctx).Save(v).Error
}
