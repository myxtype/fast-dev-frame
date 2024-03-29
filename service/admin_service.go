package service

import (
	"context"
	"errors"
	"frame/conf"
	"frame/model"
	"frame/pkg/ecode"
	"frame/pkg/jwttool"
	"frame/pkg/sql/sqltypes"
	"frame/store/db"
	"github.com/golang-jwt/jwt/v4"
	"slices"
	"time"
)

type adminService struct{}

var AdminService = new(adminService)

// Login 通过用户名密码登录
func (s *adminService) Login(ctx context.Context, username, password string) (*model.AdminUser, string, error) {
	admin, err := db.Shared().GetAdminUserByUsername(ctx, username)
	if err != nil {
		return nil, "", err
	}

	if admin == nil {
		return nil, "", errors.New("用户不存在")
	}
	if !admin.Password.Check(password) {
		return nil, "", errors.New("密码错误")
	}

	token, err := s.LoginFromAdminUser(ctx, admin)
	if err != nil {
		return nil, "", err
	}

	return admin, token, err
}

// LoginFromAdminUser 登录
func (s *adminService) LoginFromAdminUser(ctx context.Context, admin *model.AdminUser) (string, error) {
	if admin.Disabled {
		return "", errors.New("账号已被禁用")
	}

	// 更新登录版本
	admin.LoginVersion++
	if err := db.Shared().UpdateAdminUserLoginVersion(ctx, admin); err != nil {
		return "", err
	}

	token, err := s.CreateToken(admin)
	if err != nil {
		return "", err
	}

	return token, nil
}

// CreateToken 生成Token
func (s *adminService) CreateToken(user *model.AdminUser) (string, error) {
	claim := jwttool.BuildAdminClaims(user.ID, user.Password.Hash, user.LoginVersion, time.Hour*24*30)
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claim)
	return token.SignedString([]byte(conf.Get().JwtSecret.Rest))
}

// ParseToken 解析Token
func (s *adminService) ParseToken(tokenStr string) (*jwttool.AdminClaims, error) {
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

// CheckToken 检查Token
func (s *adminService) CheckToken(ctx context.Context, tokenStr string) (*model.AdminUser, error) {
	claims, err := s.ParseToken(tokenStr)
	if err != nil {
		return nil, err
	}

	// 获取用户
	user, err := db.Shared().GetAdminUser(ctx, claims.UID)
	if err != nil {
		return nil, err
	}

	// 判断是否禁用此用户
	if user == nil || user.Disabled {
		return nil, errors.New("账号已被禁用")
	}
	// 多设备登录
	if user.LoginVersion != claims.Version {
		return nil, errors.New("账号已被其他设备登录")
	}
	// 登录密码
	if user.Password.Hash != claims.PasswordHash {
		return nil, errors.New("账号密码已被修改")
	}

	return user, nil
}

// AddLog 添加管理员操作日志
func (s *adminService) AddLog(ctx context.Context, adminId uint, notes, ip string) error {
	log := &model.AdminLog{
		AdminId: adminId,
		Notes:   notes,
		Ip:      ip,
	}

	return db.Shared().AddAdminLog(ctx, log)
}

// UpdatePassword 修改密码
func (s *adminService) UpdatePassword(ctx context.Context, admin *model.AdminUser, newPass string) error {
	pass := sqltypes.NewPassword(newPass)
	admin.Password = pass

	return db.Shared().SaveAdminUser(ctx, admin)
}

// CheckAdminRole 检查权限
func (s *adminService) CheckAdminRole(ctx context.Context, admin *model.AdminUser, permit string) (bool, error) {
	role, err := db.Shared().GetAdminRole(ctx, admin.RoleId)
	if err != nil {
		return false, err
	}

	if role == nil || role.Disabled {
		return false, nil
	}

	// 全部权限
	if slices.Contains(role.Permissions, "*") {
		return true, nil
	}

	return slices.Contains(role.Permissions, permit), nil
}

// GetAdminUser 获取管理员
func (s *adminService) GetAdminUser(ctx context.Context, id uint) (*model.AdminUser, error) {
	return db.Shared().GetAdminUser(ctx, id)
}

// SaveAdminUser 保存管理员
func (s *adminService) SaveAdminUser(ctx context.Context, admin *model.AdminUser) error {
	return db.Shared().SaveAdminUser(ctx, admin)
}

// QueryAdminUsers 查询管理员
func (s *adminService) QueryAdminUsers(ctx context.Context, id, roleId int64, username string, page, limit int) ([]*model.AdminUser, int64, error) {
	return db.Shared().QueryAdminUsers(ctx, id, roleId, username, page, limit)
}

// GetAdminUserCount 获取管理员数据
func (s *adminService) GetAdminUserCount(ctx context.Context) (int64, error) {
	return db.Shared().GetAdminUserCount(ctx)
}

// QueryAdminRoles 查询管理员角色
func (s *adminService) QueryAdminRoles(ctx context.Context, page, limit int) ([]*model.AdminRole, int64, error) {
	return db.Shared().QueryAdminRoles(ctx, page, limit)
}

// GetAdminRole 获取管理员角色
func (s *adminService) GetAdminRole(ctx context.Context, id uint) (*model.AdminRole, error) {
	return db.Shared().GetAdminRole(ctx, id)
}

// SaveAdminRole 保存管理员角色
func (s *adminService) SaveAdminRole(ctx context.Context, v *model.AdminRole) error {
	return db.Shared().SaveAdminRole(ctx, v)
}

// InitAdmin 初始化管理员
func (s *adminService) InitAdmin(ctx context.Context) error {
	tx, err := db.Shared().BeginTx()
	if err != nil {
		return err
	}
	defer tx.Rollback()

	role := &model.AdminRole{
		Name:        "超级管理员",
		Permissions: []string{"*"},
	}

	if err := tx.SaveAdminRole(ctx, role); err != nil {
		return err
	}

	if err := tx.SaveAdminUser(ctx, &model.AdminUser{
		Username: "admin",
		Password: sqltypes.NewPassword("123456"),
		RoleId:   role.ID,
		Name:     "超级管理员",
	}); err != nil {
		return err
	}

	return tx.Commit()
}
