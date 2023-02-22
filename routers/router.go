package routers

import (
	"github.com/gin-gonic/gin"
)

func CreateAppRouter() *gin.Engine {
	r := gin.New()
	r.Use(gin.Logger())
	r.Use(gin.Recovery())
	r.GET("/", hello)
	return r
}

func hello(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "Hello starter Gin",
	})
}
