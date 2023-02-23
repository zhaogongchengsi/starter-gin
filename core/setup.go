package core

import (
	"context"
	"fmt"
	"os/signal"
	"syscall"
	"time"

	"github.com/server-gin/config"
	"github.com/server-gin/global"
	"github.com/server-gin/routers"
)

const (
	// 应用初始化
	Init int = 1
	// 应用启动
	Start int = 2
	// 应用关闭
	Close int = 3
	/// 应用重启
	Restart int = 4
)

type Signal struct {
	SignalType int
	Content    config.Config
}

func SetUp() {

	quit := make(chan Signal, 1)
	ctx, stop := signal.NotifyContext(context.Background(), syscall.SIGINT, syscall.SIGTERM)
	defer stop()

	db, err := CreateAppDataBase(global.DbConfig)

	global.Db = db
	if err != nil {
		panic(err)
	}

	routers := routers.CreateAppRouter()
	srv := CreateAppServer(routers, global.ServerConfig.Port, global.ServerConfig.Mode)

	close := func() {
		ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
		defer cancel()
		if err := srv.Shutdown(ctx); err != nil {
			fmt.Printf("Server forced to shutdown: %v", err)
		}
	}

	start := func() {
		fmt.Printf("\nThe service started successfully, the address is http://localhost:%d\n", global.ServerConfig.Port)
		if err := srv.ListenAndServe(); err != nil {
			fmt.Printf("Service startup failed %v\n", err)
		}
	}

	select {
	case sig := <-quit:
		if sig == Init {
			err := global.InitGlobalValues(quit)
			if err != nil {
				panic(err)
			}
		}
		if sig == Start {
			go start()
		}
		if sig == Close {
			close()
		}
	case <-ctx.Done():
		close()
	default:
	}

}
