package system

import (
	"github.com/gin-gonic/gin"
	"github.com/server-gin/common"
	"github.com/server-gin/core"
	"github.com/server-gin/global"
)

func Login(ctx *gin.Context) {
	common.NewOkResponse().Send(ctx)
}

func Captcha(ctx *gin.Context) {

	_, err := core.ConnectRedisServer(global.RedisConfig)

	if err != nil {
		common.NesFailResponse(300).ErrorToString(err).Send(ctx)
		return
	}

	common.NewOkResponse().Send(ctx)
}
