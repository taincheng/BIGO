package internal

import (
	"server/config"
	"time"

	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"gorm.io/gorm/schema"
)

var Gorm = new(_gorm)

type _gorm struct{}

// Config gorm 自定义配置
func (g *_gorm) Config(general config.GeneralDB) *gorm.Config {
	return &gorm.Config{
		Logger: logger.New(NewWriter(general), logger.Config{
			// 慢查询阈值，SQL 执行时间超过 200ms 会被标记为慢查询并记录
			SlowThreshold: 200 * time.Millisecond,
			LogLevel:      general.LogLevel(),
			Colorful:      true,
		}),
		// 命名策略
		NamingStrategy: schema.NamingStrategy{
			// 表名前缀: 例如前缀为 "sys_"，则 User 模型会自动映射到 sys_users 表
			TablePrefix: general.Prefix,
			// 是否使用单数表名
			SingularTable: general.Singular,
		},
		// 禁用外键约束, 在执行数据库迁移（AutoMigrate）时，不创建外键约束
		DisableForeignKeyConstraintWhenMigrating: true,
	}
}
