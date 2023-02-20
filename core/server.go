package core

import (
	"fmt"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/server-gin/global"
	"github.com/server-gin/routers"
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

func CreateAppServer() {
	gin.SetMode(global.Server.Mode)
	// 初始化路由
	router := routers.CreateAppRouter()

	// 设置静态文件目录
	router.Static("/static", global.Server.Static)

	server := initServer(fmt.Sprintf(":%d", global.Server.Port), router)

	fmt.Printf("\nThe service started successfully, the address is http://localhost:%d\n", global.Server.Port)

	if err := server.ListenAndServe(); err != nil {
		fmt.Printf("Service startup failed %v\n", err)
	}

}
