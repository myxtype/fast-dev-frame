package mysql

import (
	"frame/models"
	"gorm.io/gorm"
)

// 插入新用户
func (s *Store) AddUser(user *models.User) error {
	return s.db.Create(user).Error
}

// 更新用户
func (s *Store) UpdateUser(user *models.User) error {
	return s.db.Save(user).Error
}

// 根据用户ID查询用户
func (s *Store) GetUserById(id uint64) (*models.User, error) {
	var user models.User
	err := s.db.First(&user, id).Error
	if err == gorm.ErrRecordNotFound {
		return nil, nil
	}
	return &user, err
}
