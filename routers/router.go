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

	r.Static(global.ServerConfig.StaticName, global.ServerConfig.Static)
	r.StaticFile("/", global.ServerConfig.IndexHtml)

	// r.Static("/favicon.ico", "./dist/favicon.ico")

	v1 := r.Group(fmt.Sprintf("/%s", global.ServerConfig.Prefix))

	v1.GET("hello", hello)

	return r
}

func hello(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "Hello starter Gin",
	})
}
