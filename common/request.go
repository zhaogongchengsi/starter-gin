package common

import "github.com/gin-gonic/gin"

func ShouldBind[Response any](ree Response, c *gin.Context) (res *Response, err error) {
	err = c.ShouldBindJSON(ree)
	if err != nil {
		return res, err
	}
	return res, nil
}
