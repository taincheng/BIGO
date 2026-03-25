package initialize

import (
	"server/config"
	"server/global"
	"server/initialize/internal"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

// GormMysql 初始化Mysql
func GormMysql() *gorm.DB {
	mysqlConfig := global.BIGO_CONFIG.Mysql
	return initMysqlDatabase(mysqlConfig)
}

func initMysqlDatabase(m config.Mysql) *gorm.DB {
	if m.Dbname == "" {
		return nil
	}

	mysqlConfig := mysql.Config{
		DSN:                       m.Dsn(), // DSN data source name
		DefaultStringSize:         191,     // string 类型字段的默认长度，设置为 191 是因为 MySQL InnoDB 引擎在 UTF8MB4 编码下，索引的最大长度为 767 字节
		SkipInitializeWithVersion: false,   // 根据版本自动配置
	}
	// 数据库配置
	general := m.GeneralDB
	if db, err := gorm.Open(mysql.New(mysqlConfig), internal.Gorm.Config(general)); err != nil {
		panic(err)
	} else {
		// GORM 预定义的键名，执行 DDL 操作时传递数据库特定的选项，此处用来设置表的默认引擎
		db.InstanceSet("gorm:table_options", "ENGINE="+m.Engine)
		sqlDB, _ := db.DB()
		sqlDB.SetMaxIdleConns(m.MaxIdleConns)
		sqlDB.SetMaxOpenConns(m.MaxOpenConns)
		return db
	}
}
