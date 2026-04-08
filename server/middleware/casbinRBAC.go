package middleware

import (
	"server/global"
	"server/model/common/response"
	"server/utils"
	"strconv"
	"strings"

	"github.com/gin-gonic/gin"
)

func CasbinHandler() gin.HandlerFunc {
	return func(c *gin.Context) {
		claims, _ := utils.GetClaims(c)
		path := c.Request.URL.Path
		// 清理路由前缀
		obj := strings.TrimPrefix(path, global.BIGO_CONFIG.System.RouterPrefix)
		// 获取请求方法
		act := c.Request.Method
		// 获取用户的角色
		sub := strconv.Itoa(int(claims.AuthorityId))
		// 判断策略是否存在
		casbin := utils.GetCasbin()
		success, _ := casbin.Enforce(sub, obj, act)
		if !success {
			response.FailWithMessage("权限不足", c)
			c.Abort()
			return
		}
		c.Next()
	}
}
