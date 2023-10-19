package config

import (
	"github.com/spf13/viper"
)

type Configuration struct {
	ServerPort  string
	StorageDirs []string
}

func LoadConfig(path string) (Configuration, error) {
	viper.SetConfigFile(path)

	if err := viper.ReadInConfig(); err != nil {
		return Configuration{}, err
	}

	var config Configuration
	if err := viper.Unmarshal(&config); err != nil {
		return Configuration{}, err
	}

	return config, nil
}
