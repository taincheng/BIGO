package global

import (
	"server/config"

	"github.com/spf13/viper"
)

var (
	BIGO_CONFIG config.Config
	BIGO_VIPER  *viper.Viper
)
