package utils

import (
	"net"
	"server/global"
	systemReq "server/model/system/request"
	"time"

	"github.com/gin-gonic/gin"
)

func SetToken(c *gin.Context, token string, maxAge int) {
	host, _, err := net.SplitHostPort(c.Request.Host)
	if err != nil {
		// 没有端口时直接用 host
		host = c.Request.Host
	}
	if net.ParseIP(host) != nil {
		// host 是 IP（如 127.0.0.1、192.168.1.1）
		// 设置 Domain 为空字符串，Cookie 只在该 IP 下生效
		c.SetCookie(global.TokenKey, token, maxAge, "/", "", false, false)
	} else {
		// host 是域名（如 api.example.com）
		// 设置 Domain 为 host，Cookie 在该域名及其子域名下生效
		c.SetCookie(global.TokenKey, token, maxAge, "/", host, false, false)
	}
}
func GetToken(c *gin.Context) string {
	// 优先从请求头中获取 token
	token := c.Request.Header.Get(global.TokenKey)
	// 如果请求头中没有 token，则从 cookie 中获取
	if token == "" {
		jwt := NewJWT()
		// 从 cookie 中获取 token 字符串
		token, _ = c.Cookie(global.TokenKey)
		if token == "" {
			// 说明 cookie 中也没有 token 字符串，直接返回空字符串
			global.BIGO_LOG.Error("请检查 cookie 和请求头中的 token 字段是否存在")
			return token
		}
		// 解析 token，成功则更新 cookie 中过期时间
		claims, err := jwt.ParseToken(token)
		if err != nil {
			global.BIGO_LOG.Error(err.Error())
			// 解析 token 失败，返回的是空字符串
			return token
		}
		// 更新 cookie 中过期时间为剩余的时间
		SetToken(c, token, int(claims.ExpiresAt.Unix()-time.Now().Unix()))
	}
	return token
}

// GetClaims 获取claims
func GetClaims(c *gin.Context) (*systemReq.CustomClaims, error) {
	token := GetToken(c)
	jwt := NewJWT()
	claims, err := jwt.ParseToken(token)
	if err != nil {
		global.BIGO_LOG.Error(err.Error())
		return nil, err
	}
	return claims, nil
}
