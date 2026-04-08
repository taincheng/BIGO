package system

import (
	"github.com/gin-gonic/gin"
	"github.com/mojocn/base64Captcha"
)

var store = base64Captcha.DefaultMemStore

func (b *BaseApi) Captcha(c *gin.Context) {

}
