package ecode

var (
	errServerCode   = -1 // 服务器错误代码
	Ok              = add(0, "ok")
	ErrRequest      = add(400, "请求参数错误")
	ErrNotFind      = add(404, "数据异常或未找到")
	ErrForbidden    = add(403, "请求被拒绝")
	ErrNoPermission = add(405, "请求无权限")
	ErrAuth         = add(600, "认证错误")
)
