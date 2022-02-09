package conf

import (
	"github.com/spf13/viper"
)

// AtlasConfig exposed
var AtlasConfig *AtlasConfigStruct

// AtlasConfigStruct config struct
type AtlasConfigStruct struct {
	VAR1 string `mapstructure:"VAR1"`
}

func init() {
	AtlasConfig = &AtlasConfigStruct{}
}

// Config export
func Config() error {
	return viper.Unmarshal(AtlasConfig)
}
