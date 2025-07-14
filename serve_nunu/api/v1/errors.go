package v1

var (
	// common errors
	ErrSuccess             = newError(0, "ok")
	ErrBadRequest          = newError(400, "Bad Request")
	ErrUnauthorized        = newError(401, "Unauthorized")
	ErrNotFound            = newError(404, "Not Found")
	ErrInternalServerError = newError(500, "Internal Server Error")

	// more biz errors

	ErrIdNotFound        = newError(1002, "id未找到")
	ErrIdDeleted         = newError(1003, "id已删除")
	ErrTagNntFound       = newError(1004, "标签不存在")
	ErrTagDefinded       = newError(1005, "标签已定义")
	ErrUserNotFound      = newError(1006, "用户不存在")
	ErrAccountAlreadyUse = newError(1007, "账号已被使用")
	ErrTagAlreadyUse     = newError(1008, "标签已被使用")
	ErrUserAlreadyUse    = newError(1009, "用户已被使用")
)
