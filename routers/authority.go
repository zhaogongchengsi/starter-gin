package routers

import (
	"github.com/gin-gonic/gin"
	apiSystem "github.com/server-gin/api/v1/system"
	"github.com/server-gin/middleware"
)

func RegisterAuthorityRouter(group *gin.RouterGroup) {
	authRouter := group.Group("auth").Use(middleware.JWTAuth())
	{
		authRouter.GET("auth_list", apiSystem.GetAuthority)
	}
}
