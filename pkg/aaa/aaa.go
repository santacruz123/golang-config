package aaa

import (
	"fmt"

	"github.com/spf13/viper"
)

//Print func
func Print() string {

	v := fmt.Sprintf(
		"from pkg/aaa/aaa.go: VAR1=%s VAR2=%s\n",
		viper.GetString("VAR1"),
		viper.GetString("VAR2"),
	)

	fmt.Println(v)
	return v
}
