package core

import (
	"fmt"
	"server/global"
	"server/initialize"
	"time"
)

func RunServer() {

	Router := initialize.Routers()

	address := fmt.Sprintf(":%d", global.BIGO_CONFIG.System.Addr)

	fmt.Printf(`
		欢迎使用 BIGO
		默认自动化文档地址:http://127.0.0.1%s/swagger/index.html
		默认前端文件运行地址:http://127.0.0.1:8080
	`, address)
	initServer(address, Router, 10*time.Minute, 10*time.Minute)
}
