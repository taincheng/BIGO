package initialize

import (
	"net/http"
	"server/docs"
	"server/global"
	"server/middleware"
	"server/router"

	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

// Routers 初始化路由
func Routers() *gin.Engine {
	Router := gin.New()

	docs.SwaggerInfo.BasePath = global.BIGO_CONFIG.System.RouterPrefix
	Router.GET(global.BIGO_CONFIG.System.RouterPrefix+"/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	systemRouter := router.RouterGroupApp.System

	// 无需鉴权的的路由组
	PublicGroup := Router.Group(global.BIGO_CONFIG.System.RouterPrefix)
	// 需要鉴权的路由组
	PrivateGroup := Router.Group(global.BIGO_CONFIG.System.RouterPrefix)

	PrivateGroup.Use(middleware.JWTAuth()).Use(middleware.CasbinHandler())

	{
		// 健康检测
		PublicGroup.GET("/health", func(c *gin.Context) {
			c.JSON(http.StatusOK, "ok")
		})
	}

	{
		systemRouter.InitBaseRouter(PublicGroup)
	}

	{
		systemRouter.InitUserRouter(PrivateGroup)
		systemRouter.InitAuthorityRouter(PrivateGroup)
	}
	global.BIGO_ROUTER = Router.Routes()
	global.BIGO_LOG.Info("router register success")
	return Router
}
