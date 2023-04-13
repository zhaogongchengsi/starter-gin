package main

import (
	"flag"
	"github.com/zhaogongchengsi/starter-gin/config"
	"github.com/zhaogongchengsi/starter-gin/core"
	"github.com/zhaogongchengsi/starter-gin/global"
	"github.com/zhaogongchengsi/starter-gin/routers"
	"github.com/zhaogongchengsi/starter-gin/utils"
	"go.uber.org/zap"
)

//go:generate go env -w GO111MODULE=on
//go:generate go env -w GOPROXY=https://goproxy.cn,direct
//go:generate go mod tidy
//go:generate go mod download

func main() {
	var configFile string
	flag.StringVar(&configFile, "c", "./config.yaml", "the configuration file")
	flag.Parse()

	utils.Info("开始读取配置文件 %s\n", configFile)

	serverConfig, err := config.LoadServerConfig(configFile)
	if err != nil {
		panic(err)
	}

	utils.Info("✔ 读取配置文件成功 %s\n", configFile)

	global.AppConfig = serverConfig

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
