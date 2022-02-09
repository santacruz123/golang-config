package main

import (
	"fmt"
	"net/http"

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

func hello(w http.ResponseWriter, req *http.Request) {
	fmt.Fprintf(w, "from bin/cli/cli.go: VAR1='%s' VAR2='%s'\n", viper.GetString("VAR1"), viper.GetString("VAR2"))
}

func main() {
	http.HandleFunc("/", hello)
	fmt.Println("listening at http://localhost:8080")
	http.ListenAndServe(":8080", nil)
}
