package initialize

import (
	"os"
	"server/global"

	"go.uber.org/zap"
	"gorm.io/gorm"
)

func Gorm() *gorm.DB {
	switch global.BIGO_CONFIG.System.DbType {
	case "mysql":
		global.BIGO_ACTIVE_DBNAME = &global.BIGO_CONFIG.Mysql.Dbname
		return GormMysql()
	default:
		global.BIGO_ACTIVE_DBNAME = &global.BIGO_CONFIG.Mysql.Dbname
		return GormMysql()
	}
}

func RegisterTables() {
	if global.BIGO_CONFIG.System.DisableAutoMigrate {
		global.BIGO_LOG.Info("auto-migrate is disabled, skipping table registration")
		return
	}

	db := global.BIGO_DB
	// 迁移系统必备表
	err := db.AutoMigrate()
	if err != nil {
		global.BIGO_LOG.Error("register table failed", zap.Error(err))
		os.Exit(0)
	}

	err = bizModel()

	if err != nil {
		global.BIGO_LOG.Error("register biz_table failed", zap.Error(err))
		os.Exit(0)
	}
	global.BIGO_LOG.Info("register table success")
}
