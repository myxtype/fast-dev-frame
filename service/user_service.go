package service

import (
	"frame/model"
	"frame/pkg/sql/types"
	"frame/store/mysql"
)

type userService struct{}

var UserService = new(userService)

// 获取用户
func (s *userService) GetUserById(id int64) (*model.User, error) {
	return mysql.Shared().GetUserById(id)
}

// 注册
func (s *userService) Register(username, password string) error {
	user := &model.User{
		Username: username,
		Password: types.NewPassword(password),
	}

	return mysql.Shared().AddUser(user)
}
