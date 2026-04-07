package middleware

import "github.com/gin-gonic/gin"

func CasbinHandler() gin.HandlerFunc {
	return func(c *gin.Context) {

		c.Next()
	}
}
