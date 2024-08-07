package viper

import (
	"fmt"
	"github.com/spf13/viper"
)

const Config_Type = "yaml"

func GetViper(name, path string) {
	viper.SetConfigName(name)
	viper.SetConfigType(Config_Type)
	viper.AddConfigPath(path)
	err := viper.ReadInConfig()
	if err != nil {
		panic(fmt.Errorf("fatal error config file: %w", err))
	}
}
