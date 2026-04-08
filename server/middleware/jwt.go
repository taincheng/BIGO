package middleware

import (
	"server/global"
	"server/model/common/response"
	"server/utils"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
)

func JWTAuth() gin.HandlerFunc {
	return func(c *gin.Context) {
		token := utils.GetToken(c)
		if token == "" {
			response.NoAuth("未登录或非法访问，请登录", c)
			c.Abort()
			return
		}
		j := utils.NewJWT()
		// parseToken 解析 token
		claims, err := j.ParseToken(token)
		// token 过期或者无效等错误
		if err != nil {
			response.NoAuth(err.Error(), c)
			utils.ClearToken(c)
			c.Abort()
			return
		}
		c.Set(global.ClaimsKey, claims)
		// 距离过期时间还剩多少秒
		remainder := claims.ExpiresAt.Unix() - time.Now().Unix()
		// 判断剩余时间是否小于缓冲时间
		if remainder < claims.BufferTime {
			// 在缓冲时间内，需要更新 token
			claims.ExpiresAt = jwt.NewNumericDate(time.Now().Add(time.Duration(global.BIGO_CONFIG.JWT.ExpiresTime) * time.Second))
			newToken, _ := j.CreateTokenByOldToken(token, *claims)
			newClaims, _ := j.ParseToken(newToken)
			c.Header(global.NewTokenKey, newToken)
			c.Header(global.NewExpiresAtKey, strconv.FormatInt(newClaims.ExpiresAt.Unix(), 10))
			utils.SetToken(c, newToken, int(newClaims.ExpiresAt.Unix()-time.Now().Unix()))
		}
		c.Next()

		// 兜底，防止其他 middleware 设置了新的 Token，将其更新到 Header 中
		if NewToken, exists := c.Get(global.NewTokenKey); exists {
			c.Header(global.NewTokenKey, NewToken.(string))
		}
		if NewExpiresAt, exists := c.Get(global.NewExpiresAtKey); exists {
			c.Header(global.NewExpiresAtKey, NewExpiresAt.(string))
		}
	}
}
