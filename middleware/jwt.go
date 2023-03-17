package middleware

import (
	"errors"
	"github.com/gin-gonic/gin"
	"github.com/server-gin/common"
	"github.com/server-gin/global"
	"github.com/server-gin/modules/system"
	"github.com/server-gin/utils"
	"strings"
)

func JWTAuth() gin.HandlerFunc {
	return func(c *gin.Context) {
		token := c.Request.Header.Get("Authorization")
		if len(token) < 1 {
			common.NewFailResponse().SendAfterChangeMessage("未登录或非法访问", c)
			c.Abort()
			return
		}

		token = strings.Replace(token, "Bearer ", "", 1)

		userclaims, err := utils.ParseToken[system.User](token, global.AppConfig.Jwt.SigningKey)

		if err != nil {
			if errors.Is(err, utils.ErrTokenIsNotValidPeriod) {
				common.NewFailResponse().SendAfterChangeMessage("授权已过期", c)
				c.Abort()
				return
			}
			common.NewFailResponse().AddError(err, "token 无效").Send(c)
			c.Abort()
			return
		}
		utils.ShouldBindUserWith(c, userclaims)
		c.Next()
	}
}
