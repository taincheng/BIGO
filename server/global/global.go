package global

import (
	"server/config"

	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
	"go.uber.org/zap"
	"gorm.io/gorm"
)

var (
	BIGO_CONFIG        config.Server
	BIGO_VIPER         *viper.Viper
	BIGO_LOG           *zap.Logger
	BIGO_ACTIVE_DBNAME *string
	BIGO_DB            *gorm.DB
	BIGO_ROUTER        gin.RoutesInfo
)
