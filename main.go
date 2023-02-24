package main

import (
	"github.com/server-gin/cmd"
	"github.com/server-gin/core"
	"github.com/server-gin/global"
	"github.com/server-gin/routers"
)

func init() {
	cmd.ParseServerOptions(&global.ConfigDirPath, &global.ConfigType)
	cmd.ParseDevOptions(&global.IsInit)
	cmd.Parse()

	err := global.InitGlobalValues()
	if err != nil {
		panic(err)
	}
}

//go:generate go env -w GO111MODULE=on
//go:generate go env -w GOPROXY=https://goproxy.cn,direct
//go:generate go mod tidy
//go:generate go mod download

func main() {

	db, err := core.CreateAppDataBase(global.DbConfig)
	global.Db = db

	if err != nil {
		panic(err)
	}

	routers := routers.CreateAppRouter()

	core.CreateAppServer(routers, global.ServerConfig.Port, global.ServerConfig.Mode)

}
