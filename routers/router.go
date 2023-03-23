package routers

import (
	"fmt"
	systemRouter "github.com/zhaogongchengsi/starter-gin/routers/system"

	"github.com/gin-gonic/gin"
	"github.com/zhaogongchengsi/starter-gin/config"
)

func CreateAppRouter(conf *config.Config) *gin.Engine {
	r := gin.New()
	r.Use(gin.Logger())
	r.Use(gin.Recovery())

	RegisterStaticRouter(r)

	v1 := r.Group(fmt.Sprintf("/%s", conf.Server.Prefix))

	systemRouter.RegisterBaseRouter(v1)
	systemRouter.RegisterUserRouter(v1)
	systemRouter.RegisterAuthorityRouter(v1)
	systemRouter.RegisterRouterRecordRouter(v1)

	return r
}

func RegisterStaticRouter(r *gin.Engine) {
	r.Static("/assets", "./static/assets")
	r.StaticFile("/", "./static/index.html")
	r.StaticFile("/vite.svg", "./static/vite.svg")
	// http.Handle("/", http.FileServer(http.Dir("/static")))
}
