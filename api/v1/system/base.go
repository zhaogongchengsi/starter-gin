package system

import (
	"github.com/gin-gonic/gin"
	"github.com/server-gin/common"
	"github.com/server-gin/global"
)

func Login(ctx *gin.Context) {

	err := global.ReadAppConfig()

	if err != nil {
		common.NesFailResponse(300).ErrorToString(err).Send(ctx)
		return
	}

	common.NewResponse(200, global.AppConfig, "配置读取成功").Send(ctx)
}

// var store = global.Redis 使用redis
// var store = base64Captcha.DefaultMemStore

func Captcha(ctx *gin.Context) {

	// base64Captcha.NewDriverDigit()

	common.NewOkResponse().Send(ctx)
}
