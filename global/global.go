package global

import (
	"fmt"
	"github.com/server-gin/config"
)

const (
	ConfigType string = "yaml"      // 配置文件类型
	ConfigDirPath string = "configs" // 配置文件路径
)

var (
	Server = &config.Server{}
	DbConfig = &config.DataBase{}
)

func initServer () {
	err := config.ReadConfigs(ConfigDirPath + "/server.yaml", ConfigType, "Server", &Server)
	if err != nil {
		Server = &config.Server {
			Port: 3000,
			Host: "0.0.0.0",
			Mode: "debug",
			Prefix: "api/v1",
		}
		fmt.Printf("服务配置读取失败: %v, 使用默认配置", err)
	}
}

func initDbConfig ()  {
	err := config.ReadConfigs(ConfigDirPath + "/database.yaml", ConfigType, "DataBase", &DbConfig)
	if err != nil {
		fmt.Printf("数据库配置读取失败: %v", err)
	}
}


func InitGlobalValues()  {
	initServer()
	initDbConfig()
}




