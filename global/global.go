package global

import (
	"net/http"

	"github.com/go-redis/redis/v8"
	"github.com/server-gin/config"

	// "github.com/songzhibin97/gkit/cache/local_cache"
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

	// 本地缓存
	// LocalCache local_cache.Cache
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
