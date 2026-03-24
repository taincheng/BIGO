package main

import (
	"server/core"
	"server/global"
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
}

// initSystem 初始化系统所有需要用到的组件
func initSystem() {
	// 初始化 viper, 读取配置文件
	global.BIGO_VIPER = core.Viper()
	global.BIGO_LOG = core.Zap()
}
