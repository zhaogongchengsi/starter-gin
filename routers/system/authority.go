package system

import (
	"github.com/gin-gonic/gin"
	apiSystem "github.com/zhaogongchengsi/starter-gin/api/v1/system"
	"github.com/zhaogongchengsi/starter-gin/middleware"
)

func RegisterAuthorityRouter(group *gin.RouterGroup) {
	authRouter := group.Group("auth").Use(middleware.JWTAuth())
	{
		// todo: 分页查询所有权限列表
		authRouter.GET("auth_list", apiSystem.GetAuthority)
		authRouter.POST("add_auth", apiSystem.AddAuthority)
		// todo: 删除权限
		authRouter.DELETE("delete_auto", apiSystem.DeleteAuthority)
	}
}
