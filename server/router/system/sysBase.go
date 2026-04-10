package system

import "github.com/gin-gonic/gin"

type BaseRouter struct{}

func (s *BaseRouter) InitBaseRouter(Router *gin.RouterGroup) {
	group := Router.Group("base")
	{
		group.POST("login", baseApi.Login)
		group.POST("captcha", baseApi.Captcha)
	}
}
