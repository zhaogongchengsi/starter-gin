package system

import (
	"github.com/gin-gonic/gin"
	"github.com/mojocn/base64Captcha"
	"github.com/server-gin/common"
	"github.com/server-gin/global"
)

func Login(ctx *gin.Context) {

	common.NewResponse(200, global.AppConfig, "配置读取成功").Send(ctx)
}

// var store = global.Redis 使用redis
var store = base64Captcha.DefaultMemStore

type CaptchaResponse struct {
	Id  string `json:"id"`
	Url string `json:"url"`
}

func Captcha(ctx *gin.Context) {
	config := global.AppConfig.Captcha
	driver := base64Captcha.NewDriverDigit(config.Height, config.Width, config.Length, config.MaxSkew, config.DotCount)
	cp := base64Captcha.NewCaptcha(driver, store)
	id, b64s, err := cp.Generate()

	if err != nil {
		common.NewFailResponse().ErrorToString(err).Send(ctx)
		return
	}

	common.NewResponse(200, CaptchaResponse{id, b64s}, "验证码获取成功").Send(ctx)
}
