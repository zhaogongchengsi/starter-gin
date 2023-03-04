package system

import (
	"github.com/gin-gonic/gin"
	apiSystem "github.com/server-gin/api/v1/system"
)

func RegisterBaseRouter(group *gin.RouterGroup) {
	baseRouter := group.Group("base")
	{

		baseRouter.GET("captcha", apiSystem.Captcha)
	}
	userRouter := baseRouter.Group("user")
	{
		userRouter.POST("login", apiSystem.Login)
	}
}
