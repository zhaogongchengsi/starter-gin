package core

import (
	"context"
	"fmt"
	"github.com/zhaogongchengsi/starter-gin/utils"
	"net/http"
	"os/signal"
	"syscall"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/zhaogongchengsi/starter-gin/config"
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

func CreateAppServer(conf *config.Config, router *gin.Engine) {

	server := conf.Server

	gin.SetMode(server.Mode)

	ctx, stop := signal.NotifyContext(context.Background(), syscall.SIGINT, syscall.SIGTERM)
	defer stop()

	srv := initServer(fmt.Sprintf(":%d", server.Port), router)

	if server.Https.CertFile != "" && server.Https.KeyFile != "" {
		utils.Success("\nThe service started successfully, the address is https://localhost%s\n", srv.Addr)
		go ListenAndServeTLS(srv, server.Https.CertFile, server.Https.KeyFile)
	} else {
		utils.Success("\nThe service started successfully, the address is http://localhost%s\n", srv.Addr)
		go ListenAppServe(srv)
	}

	<-ctx.Done()

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	if err := srv.Shutdown(ctx); err != nil {
		fmt.Printf("Server forced to shutdown: %v", err)
	}

}

func ListenAppServe(srv *http.Server) {
	if err := srv.ListenAndServe(); err != nil {
		fmt.Printf("Service startup failed %v\n", err)
	}
}

func ListenAndServeTLS(srv *http.Server, certFile string, keyFile string) {

	err := srv.ListenAndServeTLS(certFile, keyFile)
	if err != nil {
		fmt.Printf("Service startup failed %v\n", err)
	}
}
