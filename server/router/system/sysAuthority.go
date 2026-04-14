package system

import "github.com/gin-gonic/gin"

type AuthorityRouter struct{}

func (a *AuthorityRouter) InitAuthorityRouter(Router *gin.RouterGroup) {
	group := Router.Group("authority")
	{
		group.POST("createAuthority", authorityApi.CreateAuthority) // 创建角色
	}
}
