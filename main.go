package main

import (
	"github.com/server-gin/cmd"
)

func init() {
	err := cmd.Parse()

	if err != nil {
		panic(err)
	}

	// err := global.InitGlobalValues()
	// if err != nil {
	// 	panic(err)
	// }
}

//go:generate go env -w GO111MODULE=on
//go:generate go env -w GOPROXY=https://goproxy.cn,direct
//go:generate go mod tidy
//go:generate go mod download

func main() {

	// db, err := core.CreateAppDataBase(global.DbConfig)
	// global.Db = db

	// if err != nil {
	// 	panic(err)
	// }

	// routers := routers.CreateAppRouter()

	// core.CreateAppServer(routers, global.ServerConfig.Port, global.ServerConfig.Mode)

}
