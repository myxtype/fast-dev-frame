package db

import (
	"frame/model"
	"gorm.io/gorm"
)

func (s *Store) AddUser(user *model.User) error {
	return s.db.Create(user).Error
}

func (s *Store) UpdateUser(user *model.User) error {
	return s.db.Save(user).Error
}

func (s *Store) GetUserById(id int64) (*model.User, error) {
	var user model.User
	err := s.db.First(&user, id).Error
	if err == gorm.ErrRecordNotFound {
		return nil, nil
	}
	return &user, err
}
