package routers

import (
	"fmt"
	"github.com/zhaogongchengsi/starter-gin/global"
	"github.com/zhaogongchengsi/starter-gin/middleware"
	systemRouter "github.com/zhaogongchengsi/starter-gin/routers/system"

	"github.com/gin-gonic/gin"
	"github.com/zhaogongchengsi/starter-gin/config"
)

func CreateAppRouter(conf *config.Config, mode string) *gin.Engine {
	r := gin.New()
	gin.SetMode(mode)
	r.Use(gin.Logger())
	r.Use(gin.Recovery())

	systemRouter.RegisterStaticRouter(r)

	r.Use(middleware.Cors(global.AppConfig.Cors))

	v1 := r.Group(fmt.Sprintf("/%s", conf.Server.Prefix))
	systemRouter.RegisterBaseRouter(v1)
	systemRouter.RegisterUserRouter(v1)
	systemRouter.RegisterAuthorityRouter(v1)
	systemRouter.RegisterRouterRecordRouter(v1)

	return r
}
