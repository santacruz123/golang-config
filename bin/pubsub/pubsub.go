package main

import (
	"fmt"

	_ "github.com/joho/godotenv/autoload"

	"github.com/spf13/viper"
)

func init() {
	viper.SetConfigName("config")
	viper.AddConfigPath(".")
	viper.AutomaticEnv()
	err := viper.ReadInConfig()
	if err != nil {
		panic(fmt.Errorf("fatal error config file: %w", err))
	}
}

func main() {
	// pubsub.Handler(func(ctx context.Context, msg *pubsub.Message) {
	//
	// }))
}
