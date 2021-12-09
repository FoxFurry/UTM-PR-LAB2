package config

import (
	"github.com/spf13/viper"
	"log"
)

func LoadConfig(path string) {
	viper.SetConfigFile(path)

	if err := viper.ReadInConfig(); err != nil {
		log.Fatalf("Unable to read config: %v", err)
	}
}
