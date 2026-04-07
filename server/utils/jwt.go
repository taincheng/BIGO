package utils

import (
	"errors"
	"server/global"
	"server/model/system/request"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

type JWT struct {
	SigningKey []byte
}

var (
	TokenValid            = errors.New("未知错误")
	TokenExpired          = errors.New("token已过期")
	TokenNotValidYet      = errors.New("token尚未激活")
	TokenMalformed        = errors.New("这不是一个token")
	TokenSignatureInvalid = errors.New("无效签名")
	TokenInvalid          = errors.New("无法处理此token")
)

func NewJWT() *JWT {
	return &JWT{
		[]byte(global.BIGO_CONFIG.JWT.SigningKey),
	}
}

func (j *JWT) CreateClaims(baseClaims request.BaseClaims) request.CustomClaims {

	claims := request.CustomClaims{
		BaseClaims: baseClaims,
		// 缓冲时间，在这期间，会获取新的 token 令牌，此时一个用户将有两个令牌，前端保留一个
		BufferTime: global.BIGO_CONFIG.JWT.BufferTime,
		RegisteredClaims: jwt.RegisteredClaims{
			Audience: jwt.ClaimStrings{"BIGO"},
			// 一微秒的时间差，防止时间同步问题
			NotBefore: jwt.NewNumericDate(time.Now().Add(-1000)),
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(time.Duration(global.BIGO_CONFIG.JWT.ExpiresTime) * time.Second)),
			Issuer:    global.BIGO_CONFIG.JWT.Issuer,
		},
	}
	return claims
}

// CreateToken 创建一个token
func (j *JWT) CreateToken(claims request.CustomClaims) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString(j.SigningKey)
}

// CreateTokenByOldToken 根据旧token创建一个新token，使用乐观归并锁，防止并发重复生成多个 token
func (j *JWT) CreateTokenByOldToken(oldToken string, claims request.CustomClaims) (string, error) {
	v, err, _ := global.BIGO_SINGLEFLIGHT.Do("JWT:"+oldToken, func() (interface{}, error) {
		return j.CreateToken(claims)
	})
	return v.(string), err
}

// ParseToken 解析token
func (j *JWT) ParseToken(tokenString string) (*request.CustomClaims, error) {
	token, err := jwt.ParseWithClaims(tokenString, &request.CustomClaims{}, func(token *jwt.Token) (interface{}, error) {
		return j.SigningKey, nil
	})
	if err != nil {
		switch {
		case errors.Is(err, jwt.ErrTokenExpired):
			return nil, TokenExpired
		case errors.Is(err, jwt.ErrTokenNotValidYet):
			return nil, TokenNotValidYet
		case errors.Is(err, jwt.ErrTokenMalformed):
			return nil, TokenMalformed
		case errors.Is(err, jwt.ErrTokenSignatureInvalid):
			return nil, TokenSignatureInvalid
		default:
			return nil, TokenInvalid
		}
	}
	if token != nil {
		if claims, ok := token.Claims.(*request.CustomClaims); ok && token.Valid {
			return claims, nil
		}
	}
	return nil, TokenInvalid
}
