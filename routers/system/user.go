package system

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
		userAuthRouter.POST("changepass", apiSystem.ChangePassword) // 修改密码
		//userAuthRouter.GET("getusers", apiSystem.GetUsers)
		userAuthRouter.DELETE("deleteuser", apiSystem.DeleteUser)            // 删除用户
		userAuthRouter.GET("authoritys", apiSystem.GetUserAuthorities)       // 获取用户权限
		userRouter.POST("set_authority", apiSystem.SetUserAuthority)         // 设置用户权限
		userRouter.DELETE("delete_authority", apiSystem.DeleteUserAuthority) // 删除用户权限
		userRouter.GET("getrouters", apiSystem.GetUserRouters)               // 根据用户权限 获取用户路由
	}

}
