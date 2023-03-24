package config

import (
	"github.com/spf13/viper"
)

type Viper struct {
	viper *viper.Viper
}

func NewConfig(r string, t string, files ...string) *Viper {
	v := viper.New()
	v.SetConfigType(t)
	v.AddConfigPath(r)

	for _, file := range files {
		v.SetConfigName(file)
		//v.MergeInConfig()
	}

	return &Viper{
		viper: v,
	}
}

func (c *Viper) ReadConfigs() error {
	return c.viper.ReadInConfig()
}

func (c *Viper) Unmarshal(val any) error {
	return c.viper.Unmarshal(val)
}

func (c *Viper) ReadValue(value any, key string) error {
	return c.viper.UnmarshalKey(key, value)
}
