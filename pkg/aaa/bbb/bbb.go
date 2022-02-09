package bbb

import (
	"fmt"

	"github.com/spf13/viper"

	"goconf/pkg/conf"
)

//Print func
func Print() {
	_ = conf.Config()
	fmt.Printf("from pkg/aaa/bbb/bbb.go: VAR1=%s VAR2=%s\n", conf.AtlasConfig.VAR1, viper.GetString("VAR2"))
}
