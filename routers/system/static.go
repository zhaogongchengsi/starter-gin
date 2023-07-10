package system

import (
	"github.com/gin-gonic/gin"
	"github.com/zhaogongchengsi/starter-gin/global"
)

func RegisterStaticRouter(r *gin.Engine) {

	staticPath := global.AppConfig.Server.Static
	staticName := global.AppConfig.Server.StaticName

	r.Static(staticName, staticPath)
}
