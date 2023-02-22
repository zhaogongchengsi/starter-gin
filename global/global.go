package global

import (
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
	JwtConfig    = &config.Jwt{}
	GenConfig    = &config.Gen{}
)

var (
	Server = &http.Server{}
	Db     = &gorm.DB{}
)

func InitGlobalValues() (err error) {

	ServerConfig, err = config.ReadServerConfig()

	if err != nil {
		return err
	}

	JwtConfig, err = config.ReadJwtConfig()

	if err != nil {
		return err
	}

	DbConfig, err = config.ReadDbConfig()

	if err != nil {
		return err
	}

	GenConfig, err = config.ReadGenConfig()

	if err != nil {
		return err
	}

	return nil
}
