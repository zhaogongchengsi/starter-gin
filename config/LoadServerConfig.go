package config

import "github.com/spf13/viper"

func LoadServerConfig(file string) (c *Config, e error) {
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

	return
}
