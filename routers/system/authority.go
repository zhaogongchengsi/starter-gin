package system

import (
	"github.com/gin-gonic/gin"
	apiSystem "github.com/zhaogongchengsi/starter-gin/api/v1/system"
	"github.com/zhaogongchengsi/starter-gin/middleware"
)

func RegisterAuthorityRouter(group *gin.RouterGroup) {
	authRouter := group.Group("auth").Use(middleware.JWTAuth())
	{
		authRouter.GET("authorities", apiSystem.GetAuthorities)
		authRouter.POST("add_auth", apiSystem.AddAuthority)
		// todo: 删除权限
		authRouter.DELETE("delete_auth", apiSystem.DeleteAuthority)
	}
}
