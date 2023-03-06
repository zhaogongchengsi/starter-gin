package system

import (
	"github.com/gin-gonic/gin"
	apiSystem "github.com/server-gin/api/v1/system"
)

func RegisterUserRouter(group *gin.RouterGroup) {
	userRouter := group.Group("user")
	{
		userRouter.POST("login", apiSystem.Login)
	}
}
