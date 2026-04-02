package initialize

import (
	"github.com/gin-gonic/gin"
)

// Routers 初始化路由
func Routers() *gin.Engine {
	Router := gin.New()

	//systemRouter := router.RouterGroupApp.System

	{
		//systemRouter.InitUserRouter(Router)
	}
	return Router
}
