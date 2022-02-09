package utils

import (
	"fmt"

	"github.com/spf13/viper"
)

func init() {
	viper.SetConfigName("config")

	viper.AddConfigPath(".")
	viper.AddConfigPath("..")
	viper.AddConfigPath("../../")

	viper.AutomaticEnv()

	err := viper.ReadInConfig()
	if err != nil {
		panic(fmt.Errorf("fatal error config file: %w", err))
	}
}
