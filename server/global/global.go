package global

import (
	"server/config"

	"github.com/spf13/viper"
	"go.uber.org/zap"
)

var (
	BIGO_CONFIG config.Server
	BIGO_VIPER  *viper.Viper
	BIGO_LOG    *zap.Logger
)
