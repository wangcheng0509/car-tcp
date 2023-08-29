package ginx

import (
	"encoding/json"
	"fmt"

	"net/http"
	"strings"

	"car.tcp.consumer/ginx/errors"
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"gopkg.in/validator.v2"
)

const (
	// ReqBodyKey 请求body
	ReqBodyKey = "/req-body"
	// ResBodyKey 响应body
	ResBodyKey = "/res-body"

	// header
	authorization   = "Authorization"
	jwtPrefix       = "Bearer "
	jsonContentType = "application/json; charset=utf-8"
)

// ParseJSON 解析请求JSON
func ParseJSON(c *gin.Context, obj interface{}) error {
	if err := c.ShouldBindJSON(obj); err != nil {
		return errors.Wrap400Response(err, fmt.Sprintf("解析请求参数发生错误 - %s", err.Error()))
	}
	if err := validator.Validate(obj); err != nil {
		return errors.Wrap400Response(err, fmt.Sprintf("参数校验不通过 - %s", err.Error()))
	}

	return nil
}

// ParseQuery 解析Query参数
func ParseQuery(c *gin.Context, obj interface{}) error {
	if err := c.ShouldBindQuery(obj); err != nil {
		return errors.Wrap400Response(err, fmt.Sprintf("解析请求参数发生错误 - %s", err.Error()))
	}

	return nil
}

// ParseForm 解析Form请求
func ParseForm(c *gin.Context, obj interface{}) error {
	if err := c.ShouldBindWith(obj, binding.Form); err != nil {
		return errors.Wrap400Response(err, fmt.Sprintf("解析请求参数发生错误 - %s", err.Error()))
	}
	if err := validator.Validate(obj); err != nil {
		return errors.Wrap400Response(err, fmt.Sprintf("参数校验不通过 - %s", err.Error()))
	}

	return nil
}

// ResOK 响应OK
func ResOK(c *gin.Context) {
	// ResSuccess(c, schema.StatusResult{Status: schema.OKStatus})
	resSuccess(c, SuccessResult{
		Status: OKStatus,
		Data:   StatusResult{Status: OKStatus},
	})

}

// ResList 响应列表数据
func ResList(c *gin.Context, v interface{}) {
	resSuccess(c, SuccessResult{
		Status: OKStatus,
		Data: ListResult{
			List: v,
		},
	})
}

// ResPage 响应分页数据
func ResPage(c *gin.Context, v interface{}, pr *PaginationResult) {
	resSuccess(c, SuccessResult{
		Status: OKStatus,
		Data: ListResult{
			List:       v,
			Pagination: pr,
		},
	})
}

// ResItem 响应单条数据
func ResItem(c *gin.Context, v interface{}) {
	resSuccess(c, SuccessResult{
		Status: OKStatus,
		Data:   v,
	})
}

// ResSuccess 响应成功
func resSuccess(c *gin.Context, v interface{}) {
	resJSON(c, http.StatusOK, v)
}

// ResError 响应错误
func ResError(c *gin.Context, err error, statusCode ...int) {
	var res *errors.ResponseError

	if err != nil {
		if e, ok := err.(*errors.ResponseError); ok {
			res = e
		} else {
			res = errors.UnWrapResponse(errors.ErrInternalServer)
			res.ERR = err
		}
	} else {
		res = errors.UnWrapResponse(errors.ErrInternalServer)
	}

	if len(statusCode) > 0 {
		res.StatusCode = statusCode[0]
	}

	if err := res.ERR; err != nil {
		if res.Message == "" {
			res.Message = err.Error()
		}
	}

	if code := res.Code; code >= 400 && code < 500 {

	} else if code >= 500 {

	}

	eitem := ErrorItem{
		Status:  ErrorStatus,
		Code:    res.Code,
		Message: res.Message,
	}

	resJSON(c, res.StatusCode, ErrorResult{ErrorItem: eitem})
}

// resJSON 响应JSON数据
func resJSON(c *gin.Context, statusCode int, v interface{}) {
	buf, err := json.Marshal(v)
	if err != nil {
		panic(err)
	}
	c.Set(ResBodyKey, buf)
	c.Data(statusCode, jsonContentType, buf)
	c.Abort()
}

// GetBody Get request body
func GetBody(c *gin.Context) []byte {
	if v, ok := c.Get(ReqBodyKey); ok {
		if b, ok := v.([]byte); ok {
			return b
		}
	}
	return nil
}

// GetToken 获取token
func GetToken(c *gin.Context) string {
	var token string
	token = c.GetHeader(authorization)
	if strings.HasPrefix(token, jwtPrefix) {
		token = token[len(jwtPrefix):]
	}
	return token
}
