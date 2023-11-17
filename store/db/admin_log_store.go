package db

import (
	"context"
	"frame/model"
)

// AddAdminLog 添加管理员操作日志
func (s *Store) AddAdminLog(ctx context.Context, v *model.AdminLog) error {
	return s.db.WithContext(ctx).Create(v).Error
}
