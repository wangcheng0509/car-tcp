package errors

import (
	"net/http"

	"github.com/pkg/errors"
)

// 定义别名
var (
	New          = errors.New
	Wrap         = errors.Wrap
	Wrapf        = errors.Wrapf
	WithStack    = errors.WithStack
	WithMessage  = errors.WithMessage
	WithMessagef = errors.WithMessagef
)

// 定义错误
var (
	ErrBadRequest              = New400Response("请求参数错误")
	ErrInvalidParent           = New400Response("无效的父级节点")
	ErrNotAllowDeleteWithChild = New400Response("含有子级，不能删除")
	ErrNotAllowDelete          = New400Response("资源不允许删除")
	ErrInvalidUserName         = New400Response("无效的用户名")
	ErrInvalidPassword         = New400Response("无效的密码")
	ErrInvalidUser             = New400Response("无效的用户")
	ErrUserDisable             = New400Response("用户被禁用，请联系管理员")

	ErrPhoneRegistered = NewResponseError(422, http.StatusUnprocessableEntity, "用户已注册")

	ErrNoPerm          = NewResponseError(401, 401, "无访问权限")
	ErrInvalidToken    = NewResponseError(9999, 401, "令牌失效")
	ErrNotFound        = NewResponseError(404, 404, "资源不存在")
	ErrMethodNotAllow  = NewResponseError(405, 405, "方法不被允许")
	ErrTooManyRequests = NewResponseError(429, 429, "请求过于频繁")
	ErrInternalServer  = NewResponseError(500, 500, "服务器发生错误")
)
