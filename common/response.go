package common

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

const (
	Ok       = 200
	Error    = 500
	NotFound = 404
)

type Response[D any] struct {
	Code    int    `json:"code"`
	Data    D      `json:"data"`
	Message string `json:"message"`
}

func NewResponse[D any](code int, data D, msg string) *Response[D] {
	return &Response[D]{
		Code:    code,
		Data:    data,
		Message: msg,
	}
}

func NewResponseWithData[D any](data D) *Response[D] {
	return NewResponse(Ok, data, "操作成果")
}

func NewOkResponse() *Response[map[string]interface{}] {
	return NewResponse(Ok, map[string]interface{}{}, "操作成功")
}

func NesFailResponse(code int) *Response[map[string]interface{}] {
	return NewResponse(Error, map[string]interface{}{}, "操作失败")
}

func (R *Response[D]) Send(c *gin.Context) {
	c.JSON(http.StatusOK, R)
}

func (R *Response[D]) ErrorToString(err error) *Response[D] {
	R.Message = err.Error()
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
