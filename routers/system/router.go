package system

import (
	"github.com/gin-gonic/gin"
	apiSystem "github.com/zhaogongchengsi/starter-gin/api/v1/system"
	"github.com/zhaogongchengsi/starter-gin/middleware"
)

func RegisterRouterRecordRouter(group *gin.RouterGroup) {
	router := group.Group("router").Use(middleware.JWTAuth())
	{
		router.GET("routers", apiSystem.GetRouters)
		router.POST("create_router", apiSystem.CreateRouter)
		router.DELETE("delete_router", apiSystem.DeleteRouter)
	}
}
