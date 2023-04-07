package common

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

const (
	AppId = "Starter-Gin-App"
	// Ok 执行成功
	Ok = 1
	// Failed 执行失败
	Failed = 0
	// ParamsErrorCode 参数错误
	ParamsErrorCode = 2
	// ServerError 服务器内部错误
	ServerError = 500
	// NotFound 资源找不到
	NotFound = 404
	// AuthFailed 认证失败
	AuthFailed = 401
	// AccreditFail 授权失败
	AccreditFail = 403
)

type Response[D any] struct {
	Code int `json:"code"`
	Data D   `json:"data"`

	// 给普通人的错误提示
	Message string `json:"message"`
	// 给程序员看的错误
	Error string `json:"error"`
}

func NewResponse[D any](code int, data D, msg string) *Response[D] {
	return &Response[D]{
		Code:    code,
		Data:    data,
		Message: msg,
		Error:   "",
	}
}

func NewResponseWithData[D any](data D) *Response[D] {
	return NewResponse(Ok, data, "操作成功")
}

func NewOkResponse() *Response[map[string]interface{}] {
	return NewResponse(Ok, map[string]interface{}{}, "操作成功")
}

func NewFailResponse() *Response[map[string]interface{}] {
	return NewResponse(Failed, map[string]interface{}{}, "操作失败")
}

func NewParamsError(c *gin.Context, err error) {
	res := Response[map[string]interface{}]{
		Code: ParamsErrorCode,
		Data: map[string]interface{}{},
	}
	res.Error = fmt.Sprintf(`[%s] [%s] %s`, AppId, c.FullPath(), err.Error())
	res.Message = "参数错误或者缺失"
	res.Send(c)
}

func (R *Response[D]) Send(c *gin.Context) {
	c.JSON(http.StatusOK, R)
}

func (R *Response[D]) AddError(err error, msg string) *Response[D] {

	R.Error = fmt.Sprintf(`[%s] %s`, AppId, err.Error())
	R.Message = msg
	return R
}

func (R *Response[D]) ChangeCode(code int) *Response[D] {
	R.Code = code
	return R
}

func (R *Response[D]) SendAfterChangeData(data D, c *gin.Context) {
	R.Data = data
	c.JSON(http.StatusOK, R)
}

func (R *Response[D]) SendAfterChangeMessage(msg string, c *gin.Context) {
	R.Message = msg
	c.JSON(http.StatusOK, R)
}
