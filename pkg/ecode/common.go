package ecode

var (
	Ok              = add(0, "ok")
	ErrRequest      = add(400, "请求参数错误")
	ErrNotFind      = add(404, "没有找到")
	ErrForbidden    = add(403, "请求被拒绝")
	ErrNoPermission = add(405, "无权限")
	ErrServer       = add(500, "服务器错误")
	ErrAuth         = add(600, "认证错误")
)
