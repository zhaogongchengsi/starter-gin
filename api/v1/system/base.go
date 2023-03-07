package system

import (
	"github.com/gin-gonic/gin"
	"github.com/mojocn/base64Captcha"
	"github.com/server-gin/common"
	"github.com/server-gin/global"
)

var store = base64Captcha.DefaultMemStore

type CaptchaResponse struct {
	Id  string `json:"id"`
	Url string `json:"url"`
}

func Captcha(ctx *gin.Context) {

	// 使用 redis
	// var store = core.NewRedisStore(global.Redis)

	config := global.AppConfig.Captcha
	driver := base64Captcha.NewDriverDigit(config.Height, config.Width, config.Length, config.MaxSkew, config.DotCount)
	cp := base64Captcha.NewCaptcha(driver, store)
	id, b64s, err := cp.Generate()

	if err != nil {
		common.NewParamsError(ctx, err)
		return
	}

	common.NewResponseWithData(CaptchaResponse{id, b64s}).Send(ctx)
}
