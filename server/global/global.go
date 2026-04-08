package global

import (
	"server/config"

	"github.com/gin-gonic/gin"
	"github.com/songzhibin97/gkit/cache/local_cache"
	"github.com/spf13/viper"
	"go.uber.org/zap"
	"golang.org/x/sync/singleflight"
	"gorm.io/gorm"
)

var (
	BIGO_CONFIG        config.Server
	BIGO_VIPER         *viper.Viper
	BIGO_LOG           *zap.Logger
	BIGO_ACTIVE_DBNAME *string
	BIGO_DB            *gorm.DB
	BIGO_ROUTER        gin.RoutesInfo
	BIGO_SINGLEFLIGHT  = &singleflight.Group{}
	LocalCache         local_cache.Cache
)
