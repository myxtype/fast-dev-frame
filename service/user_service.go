package service

import (
	"context"
	"frame/model"
	"frame/pkg/sql/types"
	"frame/store/db"
)

type userService struct{}

var UserService = new(userService)

// GetUserByID 获取用户
func (s *userService) GetUserByID(ctx context.Context, id int64) (*model.User, error) {
	return db.Shared().GetUserByID(ctx, id)
}

// Register 注册
func (s *userService) Register(ctx context.Context, username, password string) error {
	user := &model.User{
		Username: username,
		Password: types.NewPassword(password),
	}

	return db.Shared().AddUser(ctx, user)
}
