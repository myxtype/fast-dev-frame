package db

import (
	"context"
	"frame/model"
	"gorm.io/gorm"
)

func (s *Store) AddUser(ctx context.Context, user *model.User) error {
	return s.db.WithContext(ctx).Create(user).Error
}

func (s *Store) UpdateUser(ctx context.Context, user *model.User) error {
	return s.db.WithContext(ctx).Save(user).Error
}

func (s *Store) GetUserByID(ctx context.Context, id uint) (*model.User, error) {
	var user model.User
	err := s.db.WithContext(ctx).First(&user, id).Error
	if err == gorm.ErrRecordNotFound {
		return nil, nil
	}
	return &user, err
}

func (s *Store) GetUsersByKeys(ctx context.Context, ids []uint) ([]*model.User, error) {
	var data []*model.User
	err := s.db.WithContext(ctx).Unscoped().Where(`"id" IN (?)`, ids).Find(&data).Error
	return data, err
}
