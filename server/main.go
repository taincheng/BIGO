//go:generate swag init
package main

import (
	"server/core"
	"server/global"
	"server/initialize"
	"time"

	"github.com/songzhibin97/gkit/cache/local_cache"
	"go.uber.org/zap"
)

// @title                       BIGO Swagger API接口文档
// @version                     v1.0.0
// @description                 gin+vue的BI系统
// @securityDefinitions.apikey  ApiKeyAuth
// @in                          header
// @name                        x-token
// @BasePath                    /
func main() {
	// 初始化服务
	initSystem()
	// 运行服务
	core.RunServer()
}

// initSystem 初始化系统所有需要用到的组件
func initSystem() {
	// 初始化 viper, 读取配置文件
	global.BIGO_VIPER = core.Viper()
	// 初始化 zap，配置日志
	global.BIGO_LOG = core.Zap()
	global.LocalCache = local_cache.NewCache(
		local_cache.SetDefaultExpire(time.Duration(global.BIGO_CONFIG.JWT.ExpiresTime) * time.Second),
	)
	// zap 提供的线程安全的方式
	zap.ReplaceGlobals(global.BIGO_LOG)
	global.BIGO_DB = initialize.Gorm()
	if global.BIGO_DB != nil {
		initialize.RegisterTables() // 初始化表
	}
}
