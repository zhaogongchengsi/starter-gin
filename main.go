package main

import (
	"github.com/server-gin/cmd"
	"github.com/server-gin/core"
	"github.com/server-gin/global"
)

func init() {
	cmd.ParseServerOptions(&global.ConfigDirPath, &global.ConfigType)
	cmd.ParseDevOptions(&global.IsInit)
	cmd.Parse()
	// 初始化全局配置变量
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
	core.SetUp()
}
