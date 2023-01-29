package routers

import (
	"github.com/gin-gonic/gin"
)

func CreateAppRouter () *gin.Engine  {
	r := gin.New()
	r.Use(gin.Logger())
	r.Use(gin.Recovery())
	return r
}

