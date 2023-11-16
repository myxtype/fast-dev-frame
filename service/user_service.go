package service

import (
	"context"
	"frame/conf"
	"frame/model"
	"frame/pkg/ecode"
	"frame/pkg/jwttool"
	"frame/pkg/sql/sqltypes"
	"frame/store/db"
	"github.com/golang-jwt/jwt/v4"
	"time"
)

type userService struct{}

var UserService = new(userService)

func (s *userService) CreateToken(user *model.AdminUser) (string, error) {
	claim := jwttool.BuildAdminClaims(user.ID, user.Password.Hash, user.LoginVersion, time.Hour*24*30)
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claim)
	return token.SignedString([]byte(conf.Get().JwtSecret.Rest))
}

func (s *userService) ParseToken(tokenStr string) (*jwttool.AdminClaims, error) {
	var claim jwttool.AdminClaims

	token, err := jwt.ParseWithClaims(tokenStr, &claim, func(token *jwt.Token) (interface{}, error) {
		return []byte(conf.Get().JwtSecret.Rest), nil
	})
	if err != nil {
		return nil, ecode.ErrAuth
	}
	if !token.Valid {
		return nil, ecode.ErrAuth
	}

	return &claim, nil
}

func (s *userService) CheckToken(ctx context.Context, tokenStr string) (*model.User, error) {
	claims, err := s.ParseToken(tokenStr)
	if err != nil {
		return nil, err
	}

	// 获取用户
	user, err := db.Shared().GetUserByID(ctx, claims.UID)
	if err != nil {
		return nil, err
	}

	// 判断是否禁用此用户
	if user == nil || user.Disabled {
		return nil, ecode.ErrUserDiabled
	}

	return user, nil
}

// GetUserByID 获取用户
func (s *userService) GetUserByID(ctx context.Context, id uint) (*model.User, error) {
	return db.Shared().GetUserByID(ctx, id)
}

// Register 注册
func (s *userService) Register(ctx context.Context, username, password string) error {
	user := &model.User{
		Username: username,
		Password: sqltypes.NewPassword(password),
	}

	return db.Shared().AddUser(ctx, user)
}
