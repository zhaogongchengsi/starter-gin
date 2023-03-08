package system

import (
	"github.com/gin-gonic/gin"
	apiSystem "github.com/server-gin/api/v1/system"
)

func RegisterBaseRouter(group *gin.RouterGroup) {
	baseRouter := group.Group("base")
	{
		baseRouter.GET("captcha", apiSystem.Captcha)
		baseRouter.POST("upload", apiSystem.UpLoad)
		baseRouter.POST("upload_mult", apiSystem.UploadMult)
	}
}
