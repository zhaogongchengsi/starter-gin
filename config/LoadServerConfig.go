package config

import (
	"github.com/spf13/viper"
	"github.com/zhaogongchengsi/starter-gin/utils"
)

func LoadServerConfig(file string) (c *Config, e error) {

	utils.Info("%s -> 开始读取配置文件... \n", file)

	v := viper.New()
	v.SetConfigFile(file)

	v.SetConfigType("yaml")

	e = v.ReadInConfig()

	if e != nil {
		return nil, e
	}

	err := v.Unmarshal(&c)
	if err != nil {
		return nil, err
	}

	utils.Success("✔ 读取配置文件成功 <- %s\n", file)

	return
}
