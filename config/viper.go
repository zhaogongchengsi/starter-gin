package config

import (
	"github.com/spf13/viper"
)

type ConfingViper struct {
	viper *viper.Viper
}

func NewConfig(r string, t string, files ...string) *ConfingViper {
	v := viper.New()
	v.SetConfigType(t)
	v.AddConfigPath(r)

	for _, file := range files {
		v.SetConfigName(file)
		v.MergeInConfig()
	}

	return &ConfingViper{
		viper: v,
	}
}

func (c *ConfingViper) ReadConfigs() error {
	return c.viper.ReadInConfig()
}

func (c *ConfingViper) Unmarshal(val any) error {
	return c.viper.Unmarshal(val)
}

func (c *ConfingViper) ReadValue(valu any, key string) error {
	return c.viper.UnmarshalKey(key, valu)
}
