package request

import "frame/model"

// GetAdminUser 获取管理员用户
func (r *AppRequest) GetAdminUser() *model.AdminUser {
	v, found := r.c.Get("__admin_user")
	if found {
		return v.(*model.AdminUser)
	}
	return nil
}

// SetAdminUser 设置管理员用户
func (r *AppRequest) SetAdminUser(user *model.AdminUser) {
	r.c.Set("__admin_user", user)
}

// GetUser 获取登录用户
func (r *AppRequest) GetUser() *model.User {
	v, found := r.c.Get("___user")
	if found {
		return v.(*model.User)
	}
	return nil
}

// SetUser 设置登录用户
func (r *AppRequest) SetUser(user *model.User) {
	r.c.Set("___user", user)
}
