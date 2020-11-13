package service

import (
	"frame/models"
	"frame/models/mysql"
)

type userService struct{}

var UserService = new(userService)

// 获取用户
func (s *userService) GetUserById(id uint64) (*models.User, error) {
	return mysql.Shared().GetUserById(id)
}

// 注册
func (s *userService) Register(username, password string) error {
	user := &models.User{
		Username: username,
		Password: password,
	}

	return mysql.Shared().AddUser(user)
}
