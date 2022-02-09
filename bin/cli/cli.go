package main

import (
	"fmt"
	"goconf/pkg/aaa"
	"goconf/pkg/aaa/bbb"

	_ "github.com/joho/godotenv/autoload"

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

func main() {
	fmt.Printf("from bin/cli/cli.go: VAR1='%s' VAR2='%s'\n", viper.GetString("VAR1"), viper.GetString("VAR2"))
	bbb.Print()
	aaa.Print()
}
