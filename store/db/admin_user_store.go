package db

import (
	"context"
	"frame/model"
	"gorm.io/gorm"
)

// GetAdminUser 获取管理员
func (s *Store) GetAdminUser(ctx context.Context, id uint) (*model.AdminUser, error) {
	var admin model.AdminUser
	err := s.db.WithContext(ctx).First(&admin, id).Error
	if err == gorm.ErrRecordNotFound {
		return nil, nil
	}
	return &admin, nil
}

// GetAdminUserByUsername 通过用户名获取管理员
func (s *Store) GetAdminUserByUsername(ctx context.Context, username string) (*model.AdminUser, error) {
	var admin model.AdminUser
	err := s.db.WithContext(ctx).First(&admin, `"username" = ?`, username).Error
	if err == gorm.ErrRecordNotFound {
		return nil, nil
	}
	return &admin, nil
}

// UpdateAdminUserLoginVersion 更新管理员登录版本
func (s *Store) UpdateAdminUserLoginVersion(ctx context.Context, admin *model.AdminUser) error {
	return s.db.WithContext(ctx).Model(admin).UpdateColumn("login_version", admin.LoginVersion).Error
}

// SaveAdminUser 保存
func (s *Store) SaveAdminUser(ctx context.Context, admin *model.AdminUser) error {
	return s.db.WithContext(ctx).Save(admin).Error
}

// GetAdminUserCount 获取管理员数量
func (s *Store) GetAdminUserCount(ctx context.Context) (int64, error) {
	var count int64
	err := s.db.WithContext(ctx).Model(&model.AdminUser{}).Count(&count).Error
	return count, err
}

// QueryAdminUsers 查询管理员
func (s *Store) QueryAdminUsers(ctx context.Context, id, roleId int64, username string, page, limit int) ([]*model.AdminUser, int64, error) {
	db := s.db.WithContext(ctx)

	if id > 0 {
		db = db.Where(`"id" = ?`, id)
	}
	if roleId > 0 {
		db = db.Where(`"role_id" = ?`, roleId)
	}
	if username != "" {
		db = db.Where(`"username" = ?`, username)
	}

	var data []*model.AdminUser
	var count int64
	err := db.Offset((page - 1) * limit).Limit(limit).Order(`"id" DESC`).Find(&data).Offset(-1).Limit(-1).Count(&count).Error

	return data, count, err
}
