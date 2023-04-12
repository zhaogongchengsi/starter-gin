package global

import (
	"github.com/zhaogongchengsi/starter-gin/config"
	"github.com/zhaogongchengsi/starter-gin/core"
	"github.com/zhaogongchengsi/starter-gin/core/store"
	"github.com/zhaogongchengsi/starter-gin/utils"
	"go.uber.org/zap"
	"time"

	// "github.com/songzhibin97/gkit/cache/local_cache"
	"gorm.io/gorm"
)

var (
	ConfigType    string = "yaml" // 配置文件类型
	ConfigDirPath string = "./"   // 配置文件路径
	ConfigName    string = "config"
)

var (
	AppConfig = &config.Config{}
)

var (
	//Server           *http.Server  = nil

	Db *gorm.DB = nil

	//Redis            *redis.Client = nil

	Logger *zap.Logger = nil
	// CaptchaStore 存储验证码
	CaptchaStore *store.CaptchaBucket = store.NewCaptchaBucket()
	// Blacklist token 黑名单
	Blacklist *core.Bucket = core.NewBucket(time.Duration(10))
)

func p(ConfigDirPath, ConfigType, name string) {
	utils.Info("\n正在从 [ %s ] 读取 [ %s ] 类型的配置文件：[ %s ]\n", ConfigDirPath, ConfigType, name)
}

func ReadConfig(ConfigDirPath, ConfigType, name string) (conf *config.Config, err error) {

	p(ConfigDirPath, ConfigType, name)

	sc := config.NewConfig(ConfigDirPath, ConfigType, name)
	err = sc.ReadConfigs()
	if err != nil {
		return conf, err
	}

	err = sc.Unmarshal(&conf)

	if err != nil {
		return conf, err
	}

	utils.Info("✔ 读取配置完成\n")

	return conf, nil
}

func ReadAppConfig() error {
	p(ConfigDirPath, ConfigType, ConfigName)
	sc := config.NewConfig(ConfigDirPath, ConfigType, ConfigName)
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
