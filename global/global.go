package global

import (
	"fmt"
	"net/http"

	"github.com/server-gin/config"
	"gorm.io/gorm"
)

const (
	ConfigType    string = "yaml"    // 配置文件类型
	ConfigDirPath string = "configs" // 配置文件路径
)

var (
	ServerConfig = &config.Server{}
	DbConfig     = &config.DataBase{}
	Server       = &http.Server{}
	Db           = &gorm.DB{}
)

func initServer() error {
	err := config.ReadConfigs(ConfigDirPath+"/server.yaml", ConfigType, "Server", &ServerConfig)
	if err != nil {
		ServerConfig = &config.Server{
			Port:   3000,
			Host:   "0.0.0.0",
			Mode:   "debug",
			Prefix: "api/v1",
		}
		fmt.Printf("服务配置读取失败: %v, 使用默认配置\n", err)
		return err
	}
	return nil
}

func initDbConfig() error {
	err := config.ReadConfigs(ConfigDirPath+"/database.yaml", ConfigType, "DataBase", &DbConfig)
	if err != nil {
		fmt.Printf("数据库配置读取失败: %v\n", err)
		return err
	}
	return nil
}

func InitGlobalValues() error {
	err := initServer()

	if err != nil {
		return err
	}

	err = initDbConfig()

	return err
}
