package system

import (
	"github.com/gin-gonic/gin"
	apiSystem "github.com/server-gin/api/v1/system"
)

func RegisterBaseRouter(group *gin.RouterGroup) {
	baseRouter := group.Group("base")
	{
		baseRouter.POST("login", apiSystem.Captcha)
		baseRouter.GET("captcha", apiSystem.Captcha)
	}
}
