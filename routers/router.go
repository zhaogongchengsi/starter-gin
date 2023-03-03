package routers

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/server-gin/global"
	"github.com/server-gin/routers/system"
)

func CreateAppRouter() *gin.Engine {
	r := gin.New()
	r.Use(gin.Logger())
	r.Use(gin.Recovery())

	RegisterStaticRouter(r)

	v1 := r.Group(fmt.Sprintf("/%s", global.ServerConfig.Prefix))

	system.RegisterBaseRouter(v1)

	return r
}

func RegisterStaticRouter(r *gin.Engine) {
	r.Static("/assets", "./static/assets")
	r.StaticFile("/", "./static/index.html")
	r.StaticFile("/vite.svg", "./static/vite.svg")
	// http.Handle("/", http.FileServer(http.Dir("/static")))
}
