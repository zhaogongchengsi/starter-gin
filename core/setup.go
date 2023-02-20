package core

import (
	"fmt"

	"github.com/server-gin/global"
	"github.com/server-gin/routers"
)

func SetUp() {

	routers := routers.CreateAppRouter()
	server := CreateAppServer(routers, global.Server.Port, global.Server.Mode)

	fmt.Printf("\nThe service started successfully, the address is http://localhost:%d\n", global.Server.Port)

	if err := server.ListenAndServe(); err != nil {
		fmt.Printf("Service startup failed %v\n", err)
	}
}
