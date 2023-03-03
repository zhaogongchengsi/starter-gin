package global

import (
	"net/http"

	"github.com/go-redis/redis/v8"
	"github.com/server-gin/config"
	"gorm.io/gorm"
)

var (
	ConfigType    string = "yaml" // 配置文件类型
	ConfigDirPath string = "./"   // 配置文件路径
)

var (
	ServerConfig = &config.Server{}
	DbConfig     = &config.DataBase{}
	JwtConfig    = &config.Jwt{}
	GenConfig    = &config.Gen{}
	RedisConfig  = &config.Redis{}
	AppConfig    = &config.Config{}
)

var (
	Server = &http.Server{}
	Db     = &gorm.DB{}
	Redis  = &redis.Client{}
)

func ReadAppConfig() error {
	sc := config.NewConfig(ConfigDirPath, ConfigType, "config")
	err := sc.ReadConfigs()
	if err != nil {
		return err
	}

	err = sc.Unmarshal(&AppConfig)

	if err != nil {
		return err
	}

	return nil
}
