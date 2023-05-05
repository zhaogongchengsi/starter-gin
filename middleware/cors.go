package middleware

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type CoreOption struct {
	AllowHeaders string
	AllowMethods string
	AllCred      string
	ExpHeaders   string
}

var defaultOpt = &CoreOption{
	AllowHeaders: "Content-Type,AccessToken,X-CSRF-Token, Authorization, Token, new-token",
	AllowMethods: "POST, GET, OPTIONS,DELETE,PUT",
	ExpHeaders:   "Content-Length, Access-Control-Allow-Origin, Access-Control-Allow-Headers, Content-Type",
	AllCred:      "true",
}

func DefaultCoreOption() *CoreOption {
	return defaultOpt
}

func Cors(opt CoreOption) gin.HandlerFunc {
	return func(c *gin.Context) {
		method := c.Request.Method
		origin := c.Request.Header.Get("Origin")

		c.Header("Access-Control-Allow-Origin", origin)
		c.Header("Access-Control-Allow-Headers", opt.AllowHeaders)
		c.Header("Access-Control-Allow-Methods", opt.AllowMethods)
		c.Header("Access-Control-Expose-Headers", opt.ExpHeaders)
		c.Header("Access-Control-Allow-Credentials", opt.AllCred)

		// 放行所有OPTIONS方法
		if method == "OPTIONS" {
			c.AbortWithStatus(http.StatusNoContent)
		}
		// 处理请求
		c.Next()
	}
}
