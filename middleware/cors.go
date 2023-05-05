package middleware

import (
	"github.com/zhaogongchengsi/starter-gin/config"
	"net/http"

	"github.com/gin-gonic/gin"
)

func Cors(opt config.Cors) gin.HandlerFunc {
	return func(c *gin.Context) {
		method := c.Request.Method
		//origin := c.Request.Header.Get("Origin")

		c.Header("Access-Control-Allow-Origin", opt.Origin)
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
