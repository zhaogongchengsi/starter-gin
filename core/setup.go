package core

import (
	"fmt"

	"github.com/server-gin/global"
	"github.com/server-gin/routers"
)

func SetUp() {
	db, err := CreateAppDataBase(global.DbConfig)

	if err != nil {
		panic(err)
	}

	global.Db = db

	routers := routers.CreateAppRouter()
	global.Server = CreateAppServer(routers, global.ServerConfig.Port, global.ServerConfig.Mode)

	fmt.Printf("\nThe service started successfully, the address is http://localhost:%d\n", global.ServerConfig.Port)

	if err := global.Server.ListenAndServe(); err != nil {
		fmt.Printf("Service startup failed %v\n", err)
	}
}
