package service

import (
	"context"
	"frame/conf"
	"frame/model"
	"frame/pkg/ecode"
	"frame/pkg/jwttool"
	"frame/pkg/logger"
	"frame/pkg/sql/sqltypes"
	"frame/pkg/utils"
	"frame/store/db"
	"frame/store/rdb"
	"github.com/golang-jwt/jwt/v4"
	"time"
)

type userService struct{}

var UserService = new(userService)

func (s *userService) CreateToken(user *model.User) (string, error) {
	// 根据需求调整
	// 7天过期时间，每次请求用户信息时最好重新生成一个返回给前端
	// 这样就是用户7天没有登录，token就过期，需要重新登录
	claim := jwttool.BuildUserClaims(user.ID, time.Hour*24*7)

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claim)
	return token.SignedString([]byte(conf.Get().JwtSecret.Rest))
}

func (s *userService) ParseToken(tokenStr string) (*jwttool.UserClaims, error) {
	var claim jwttool.UserClaims

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
	user, err := s.GetUserCache(ctx, claims.UID)
	if err != nil {
		return nil, err
	}

	// 判断是否禁用此用户
	if user == nil || user.Disabled {
		return nil, ecode.ErrUserDisabled
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

// GetUsers 批量获取用户数据
func (s *userService) GetUsers(ctx context.Context, ids []uint) (map[uint]*model.User, error) {
	recordMap, err := rdb.Shared().GetUsers(ctx, ids)
	if err != nil {
		return nil, err
	}

	missingIds := utils.GetMissingIds(ids, recordMap)

	if len(missingIds) > 0 {
		missingRecords, err := db.Shared().GetUsersByKeys(ctx, missingIds)
		if err != nil {
			return nil, err
		}

		err = rdb.Shared().CacheUsers(ctx, missingRecords)
		if err != nil {
			logger.Sugar.Error(err)
		}

		for _, record := range missingRecords {
			recordMap[record.ID] = record
		}
	}

	return recordMap, nil
}

// GetUserCache 从缓存中获取用户数据
func (s *userService) GetUserCache(ctx context.Context, id uint) (*model.User, error) {
	records, err := s.GetUsers(ctx, []uint{id})
	if err != nil {
		return nil, err
	}
	if v, found := records[id]; found {
		return v, nil
	}
	return nil, nil
}
