package core

import (
	"context"
	"fmt"
	"net/http"
	"os/signal"
	"syscall"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/server-gin/config"
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

func CreateAppServer(config *config.Server, router *gin.Engine) {
	gin.SetMode(config.Mode)

	ctx, stop := signal.NotifyContext(context.Background(), syscall.SIGINT, syscall.SIGTERM)
	defer stop()

	srv := initServer(fmt.Sprintf(":%d", config.Port), router)

	if config.Https.CertFile != "" && config.Https.KeyFile != "" {
		fmt.Printf("\nThe service started successfully, the address is https://localhost:%d\n", config.Port)
		go ListenAndServeTLS(srv, config.Https.CertFile, config.Https.KeyFile)
	} else {
		fmt.Printf("\nThe service started successfully, the address is http://localhost:%d\n", config.Port)
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
