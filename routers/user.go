package routers

import (
	"github.com/gin-gonic/gin"
	apiSystem "github.com/zhaogongchengsi/starter-gin/api/v1/system"
	"github.com/zhaogongchengsi/starter-gin/middleware"
)

func RegisterUserRouter(group *gin.RouterGroup) {
	userRouter := group.Group("user")
	{
		userRouter.POST("login", apiSystem.Login)
		userRouter.POST("register", apiSystem.Register)
	}
	// 这个路由需要鉴权
	userAuthRouter := group.Group("user").Use(middleware.JWTAuth())
	{
		userAuthRouter.POST("changepass", apiSystem.ChangePassword)
		userAuthRouter.GET("getusers", apiSystem.GetUsers)
		userAuthRouter.DELETE("deleteuser", apiSystem.DeleteUser)
		userAuthRouter.GET("authoritys", apiSystem.GetAuths)
		userAuthRouter.GET("routers", apiSystem.GetUserRouters)
	}

}
