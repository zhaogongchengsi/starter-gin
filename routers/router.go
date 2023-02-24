package routers

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/server-gin/global"
)

func CreateAppRouter() *gin.Engine {
	r := gin.New()
	r.Use(gin.Logger())
	r.Use(gin.Recovery())

	r.Static("/static", global.ServerConfig.Static)

	r.GET("/", func(ctx *gin.Context) {
		ctx.Redirect(http.StatusFound, global.ServerConfig.Static+"/index.html")
	})

	v1 := r.Group(fmt.Sprintf("/%s", global.ServerConfig.Prefix))

	v1.GET("hello", hello)

	return r
}

func hello(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "Hello starter Gin",
	})
}
