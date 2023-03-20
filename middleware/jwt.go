package middleware

import (
	"errors"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v4"
	"github.com/zhaogongchengsi/starter-gin/common"
	"github.com/zhaogongchengsi/starter-gin/global"
	"github.com/zhaogongchengsi/starter-gin/module"
	"github.com/zhaogongchengsi/starter-gin/utils"
	"math"
	"strings"
	"time"
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

		claims, err := utils.ParseToken[module.User](token, global.AppConfig.Jwt.SigningKey)

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

		p := timeprogress(claims.IssuedAt.Time, claims.ExpiresAt.Time)

		// token 有效时间小于十分之一 重新签发新token
		if p < 10 {
			jwtConf := global.AppConfig.Jwt
			it := time.Now()
			et := it.Add(time.Duration(jwtConf.ExpiresAt) * time.Minute)
			user := claims.Info
			token, err = utils.CreateToken(user, jwtConf.SigningKey, it, et, jwtConf.Issuer)
			c.Header("new-authorization", token)
			c.Header("new-expires-at", et.String())
			c.Header("new-issued-at", it.String())
		}

		utils.ShouldBindUserWith[module.User](c, claims)
		c.Next()
	}
}

func getunix(date *jwt.NumericDate) int64 {
	return date.Time.Unix()
}

func timeprogress(startdate, enddate time.Time) int {
	now := time.Now().Unix()
	sut := startdate.Unix()
	eut := enddate.Unix()

	if now > eut {
		return 0
	}

	if now < sut {
		return 100
	}

	xaxised := eut - sut
	xaxising := eut - now

	scale := math.Ceil(float64((xaxised) / 100)) // 计算出刻度
	plan := math.Ceil(float64(xaxising) / scale) // 根据刻度计算出当前时间占用的进度

	return int(plan)
}
