package admin

import (
	"errors"
	"frame/model"
	"frame/pkg/ecode"
	"frame/pkg/request"
	"frame/pkg/sql/sqltypes"
	"frame/service"
	"github.com/gin-gonic/gin"
)

type AdminController struct{}

func (c *AdminController) Current(ctx *gin.Context) {
	app := request.New(ctx)

	user := app.GetAdminUser()

	app.Success(NewAdminUserInfoVo(user))
}

type AdminUpdatePasswordRequest struct {
	OldPassword string `json:"oldPassword"`
	NewPassword string `json:"newPassword"`
}

func (c *AdminController) UpdatePassword(ctx *gin.Context) {
	app := request.New(ctx)

	var req AdminUpdatePasswordRequest
	if err := ctx.Bind(&req); err != nil {
		app.Error(err)
		return
	}

	admin := app.GetAdminUser()
	if !admin.Password.Check(req.OldPassword) {
		app.Error(ecode.ErrUserPassword)
		return
	}

	app.Response(service.AdminService.UpdatePassword(ctx, admin, req.NewPassword))
}

type QueryAdminUsersRequest struct {
	PageRequest
	ID       int64  `json:"id" form:"id"`
	Username string `json:"username" form:"username"`
	RoleId   int64  `json:"roleId" form:"roleId"`
}

func (c *AdminController) QueryAdminUsers(ctx *gin.Context) {
	app := request.New(ctx)

	var req QueryAdminUsersRequest
	if err := ctx.Bind(&req); err != nil {
		app.Error(err)
		return
	}

	data, count, err := service.AdminService.QueryAdminUsers(ctx, req.ID, req.RoleId, req.Username, req.Page, req.Limit)
	if err != nil {
		app.Error(err)
		return
	}

	vos := []*AdminUserVo{}
	for _, n := range data {
		vos = append(vos, NewAdminUserVo(n))
	}

	app.Success(NewListResult(count, vos))
}

type QueryAdminRolesRequest struct {
	PageRequest
}

func (c *AdminController) QueryAdminRoles(ctx *gin.Context) {
	app := request.New(ctx)

	var req QueryAdminRolesRequest
	if err := ctx.Bind(&req); err != nil {
		app.Error(err)
		return
	}

	data, count, err := service.AdminService.QueryAdminRoles(ctx, req.Page, req.Limit)
	if err != nil {
		app.Error(err)
		return
	}

	vos := []*AdminRoleVo{}
	for _, n := range data {
		vos = append(vos, NewAdminRoleVo(n))
	}

	app.Success(NewListResult(count, vos))
}

type SaveAdminUserRequest struct {
	ID       uint   `json:"id"`
	Username string `json:"username"` // 用户名
	RoleId   uint   `json:"roleId"`   // 角色
	Name     string `json:"name"`     // 昵称
	Disabled bool   `json:"disabled"` // 是否禁用
	Password string `json:"password"` // 密码
}

// 保存管理员
func (c *AdminController) SaveAdminUser(ctx *gin.Context) {
	app := request.New(ctx)

	var req SaveAdminUserRequest
	if err := ctx.Bind(&req); err != nil {
		app.Error(err)
		return
	}

	var user *model.AdminUser

	if req.ID > 0 {
		var err error
		user, err = service.AdminService.GetAdminUser(ctx, req.ID)
		if err != nil {
			app.Error(err)
			return
		}
		if user == nil {
			app.Error(ecode.ErrNotFind)
			return
		}
	} else {
		user = &model.AdminUser{}
		if req.Password == "" {
			app.Error(errors.New("请输入登录密码"))
			return
		}
	}

	user.Name = req.Name
	user.RoleId = req.RoleId
	user.Username = req.Username
	user.Disabled = req.Disabled
	if req.Password != "" {
		tmp := sqltypes.NewPassword(req.Password)
		user.Password = tmp
	}

	app.Response(service.AdminService.SaveAdminUser(ctx, user))
}

type SaveAdminRoleRequest struct {
	ID          uint                 `json:"id"`
	Name        string               `json:"name"`        // 角色名称
	Permissions sqltypes.StringArray `json:"permissions"` // 权限列表
	Disabled    bool                 `json:"disabled"`    // 是否禁用
}

func (c *AdminController) SaveAdminRole(ctx *gin.Context) {
	app := request.New(ctx)

	var req SaveAdminRoleRequest
	if err := ctx.Bind(&req); err != nil {
		app.Error(err)
		return
	}

	var role *model.AdminRole
	if req.ID > 0 {
		var err error
		role, err = service.AdminService.GetAdminRole(ctx, req.ID)
		if err != nil {
			app.Error(err)
			return
		}
		if role == nil {
			app.Error(ecode.ErrNotFind)
			return
		}
	} else {
		role = &model.AdminRole{}
	}

	role.Name = req.Name
	role.Permissions = req.Permissions
	role.Disabled = req.Disabled

	err := service.AdminService.SaveAdminRole(ctx, role)
	app.Response(err)
}
