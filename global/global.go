package global

import (
	"net/http"

	"github.com/server-gin/config"
	"gorm.io/gorm"
)

var (
	ConfigType    string = "yaml"    // 配置文件类型
	ConfigDirPath string = "configs" // 配置文件路径
)

var (
	ServerConfig = &config.Server{}
	DbConfig     = &config.DataBase{}
	JwtConfig    = &config.Jwt{}
	GenConfig    = &config.Gen{}
	IsInit       = false
)

var (
	Server = &http.Server{}
	Db     = &gorm.DB{}
)

func InitGlobalValues() (err error) {
	sc := config.NewConfig(ConfigDirPath, ConfigType, "server")
	err = sc.ReadConfigs()
	if err != nil {
		return err
	}

	err = sc.ReadValue(&ServerConfig, "Server")
	if err != nil {
		return err
	}
	err = sc.ReadValue(&JwtConfig, "Jwt")
	if err != nil {
		return err
	}

	dc := config.NewConfig(ConfigDirPath, ConfigType, "database")

	err = dc.ReadValue(&DbConfig, "DataBase")

	if err != nil {
		return err
	}

	err = dc.ReadValue(&GenConfig, "Gen")

	if err != nil {
		return err
	}

	return nil

}
