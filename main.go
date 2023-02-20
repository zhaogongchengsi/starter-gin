package main

import (
	"github.com/server-gin/core"
	"github.com/server-gin/global"
)

func init() {
	// 初始化全局配置变量
	err := global.InitGlobalValues()
	if err != nil {
		panic(err)
	}
}

func main() {

	core.CreateAppServer()

}
