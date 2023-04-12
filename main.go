package main

import (
	"github.com/zhaogongchengsi/starter-gin/cmd"
	"github.com/zhaogongchengsi/starter-gin/core"
	"github.com/zhaogongchengsi/starter-gin/global"
	"github.com/zhaogongchengsi/starter-gin/routers"
	"go.uber.org/zap"
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

	logger, err := core.CreateLogger(global.AppConfig)

	if err != nil {
		panic(err)
	}

	global.Logger = logger

	zap.ReplaceGlobals(global.Logger)

	db, err := core.ConnectDataBaseServer(global.AppConfig)
	defer func() {
		db, _ := db.DB()
		err := db.Close()
		if err != nil {
			return
		}
	}()

	if err != nil {
		panic(err)
	}

	global.Db = db

	// 开启这句 连接 redis
	//rdb, err := core.ConnectRedisServer(global.AppConfig)
	//if err != nil {
	//	panic(err)
	//}
	//global.Redis = rdb

	r := routers.CreateAppRouter(global.AppConfig)

	core.CreateAppServer(global.AppConfig, r)

}
