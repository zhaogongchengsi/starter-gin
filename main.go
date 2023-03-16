package main

import (
	"github.com/server-gin/cmd"
	"github.com/server-gin/core"
	"github.com/server-gin/global"
	"github.com/server-gin/routers"
)

func init() {
	err := cmd.Parse()
	if err != nil {
		panic(err)
	}

	err = global.ReadAppConfig()
	if err != nil {
		panic(err)
	}
}

//go:generate go env -w GO111MODULE=on
//go:generate go env -w GOPROXY=https://goproxy.cn,direct
//go:generate go mod tidy
//go:generate go mod download

func main() {

	db, err := core.ConnectDataBaseServer(global.AppConfig)

	if err != nil {
		panic(err)
	}

	global.Db = db

	rdb, err := core.ConnectRedisServer(global.AppConfig)

	if err != nil {
		panic(err)
	}

	global.Redis = rdb

	r := routers.CreateAppRouter(global.AppConfig)

	core.CreateAppServer(global.AppConfig, r)

}
