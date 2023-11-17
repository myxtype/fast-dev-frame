package request

import "frame/model"

func (r *AppRequest) GetAdminUser() *model.AdminUser {
	v, found := r.c.Get("__admin_user")
	if found {
		return v.(*model.AdminUser)
	}
	return nil
}

func (r *AppRequest) SetAdminUser(user *model.AdminUser) {
	r.c.Set("__admin_user", user)
}

func (r *AppRequest) GetUser() *model.User {
	v, found := r.c.Get("___user")
	if found {
		return v.(*model.User)
	}
	return nil
}

func (r *AppRequest) SetUser(user *model.User) {
	r.c.Set("___user", user)
}
