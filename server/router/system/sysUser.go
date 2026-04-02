package system

import "github.com/gin-gonic/gin"

type UserRouter struct{}

func (u *UserRouter) InitUserRouter(Router *gin.RouterGroup) {
	userRouter := Router.Group("user")

	{
		userRouter.POST("admin_register") // 管理员注册账号
	}
}
