package system

import (
	"github.com/gin-gonic/gin"
	apiSystem "github.com/server-gin/api/v1/system"
)

func RegisterUserRouter(group *gin.RouterGroup) {
	userRouter := group.Group("user")
	{
		userRouter.POST("login", apiSystem.Login)
		userRouter.POST("register", apiSystem.Register)
	}
	// 这个路由需要鉴权
	userAuthRouter := group.Group("user")
	{
		userAuthRouter.POST("changepass", apiSystem.ChangePassword)
		userAuthRouter.DELETE("deleteuser", apiSystem.DeleteUser)
	}
}
