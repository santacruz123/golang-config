package bbb

import (
	"fmt"

	"github.com/spf13/viper"
)

//Print func
func Print() {
	fmt.Printf("from pkg/aaa/bbb/bbb.go: VAR1=%s VAR2=%s\n", viper.GetString("VAR1"), viper.GetString("VAR2"))
}
