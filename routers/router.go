package routers

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/server-gin/global"
)

func CreateAppRouter() *gin.Engine {
	r := gin.New()
	r.Use(gin.Logger())
	r.Use(gin.Recovery())
	CreateStaticRouter(r)

	v1 := r.Group(fmt.Sprintf("/%s", global.ServerConfig.Prefix))

	v1.GET("hello", hello)

	return r
}

func hello(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "Hello starter Gin",
	})
}

func CreateStaticRouter(r *gin.Engine) {
	r.Static("/assets", "./static/assets")
	r.StaticFile("/", "./static/index.html")
	// http.Handle("/", http.FileServer(http.Dir("/static")))
}
