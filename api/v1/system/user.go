package system

import (
	"github.com/gin-gonic/gin"
	"github.com/server-gin/common"
	systemService "github.com/server-gin/service/system"
)

type captcha struct {
	Id    string `binding:"required" json:"id"`
	Value string `binding:"required" json:"value"`
}

type LoginRes struct {
	systemService.Login
	Captcha captcha `binding:"required" json:"captcha"`
}

func Login(c *gin.Context) {
	var loginRes LoginRes
	err := c.ShouldBindJSON(&loginRes)
	if err != nil {
		common.NewFailResponse().ErrorToString(err).Send(c)
		return
	}

	// 使用 redis
	// var store = core.NewRedisStore(global.Redis)

	// if isOk := store.Verify(loginRes.Captcha.Id, loginRes.Captcha.Value, true); isOk {
	// 	common.NewFailResponse().SendAfterChangeMessage("验证码验证失败", c)
	// 	return
	// }

	login := systemService.Login{
		Phone:    loginRes.Phone,
		Password: loginRes.Password,
		NickName: loginRes.NickName,
		Email:    loginRes.Email,
	}

	user, token, err := login.Login()

	if err != nil {
		common.NewFailResponse().ErrorToString(err).Send(c)
		return
	}

	common.NewResponse(200, map[string]any{"user": user, "token": token}, "登录成功").Send(c)
}
