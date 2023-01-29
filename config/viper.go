package config

import "github.com/spf13/viper"



func InitViper(path, t string) (*viper.Viper, error)  {
	v := viper.New()
	v.SetConfigType(t)
	v.SetConfigFile(path)
	err := v.ReadInConfig()
	return v, err
}




