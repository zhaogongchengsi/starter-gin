package core

import (
	"fmt"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

func initServer(address string, router *gin.Engine) *http.Server {
	return &http.Server{
		Addr:           address,
		Handler:        router,
		ReadTimeout:    20 * time.Second,
		WriteTimeout:   20 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}
}

func CreateAppServer(router *gin.Engine, prot int, mode string) *http.Server {
	gin.SetMode(mode)
	return initServer(fmt.Sprintf(":%d", prot), router)
}
