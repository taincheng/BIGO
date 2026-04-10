package system

import (
	"server/global"
	"server/model/common/response"
	systemRes "server/model/system/response"
	"server/utils"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/mojocn/base64Captcha"
	"go.uber.org/zap"
)

var store = base64Captcha.DefaultMemStore

// Captcha
// @Tags Base
// @Summary 获取验证码
// @Security ApiKeyAuth
// @accept application/json
// @Produce  application/json
// @Success 200 {object} response.Response{data=systemRes.SysCaptchaResponse,msg=string} "获取验证码"
// @Router /base/captcha [post]
func (b *BaseApi) Captcha(c *gin.Context) {
	oc := b.judgeUseCaptcha(c)
	// 不需要返回验证码，直接返回成功
	if !oc {
		response.OkWithDetailed(systemRes.SysCaptchaResponse{
			CaptchaId:     "",
			PicPath:       "",
			CaptchaLength: 0,
			OpenCaptcha:   false,
		}, "无需验证码", c)
		return
	}
	// 验证码配置
	driverDigit := base64Captcha.NewDriverDigit(global.BIGO_CONFIG.Captcha.ImgHeight, global.BIGO_CONFIG.Captcha.ImgWidth, global.BIGO_CONFIG.Captcha.KeyLong, 0.7, 80)
	captcha := base64Captcha.NewCaptcha(driverDigit, store)
	id, b64s, _, err := captcha.Generate()
	if err != nil {
		global.BIGO_LOG.Error("验证码获取失败", zap.Error(err))
		response.FailWithMessage("验证码获取失败", c)
		return
	}

	response.OkWithDetailed(systemRes.SysCaptchaResponse{
		CaptchaId:     id,
		PicPath:       b64s,
		CaptchaLength: global.BIGO_CONFIG.Captcha.KeyLong,
		OpenCaptcha:   oc,
	}, "获取验证码成功", c)
}

// judgeUseCaptcha 判断是否需要校验验证码
func (b *BaseApi) judgeUseCaptcha(c *gin.Context) (oc bool) {
	clientIP := c.ClientIP()
	// 验证码开关，防爆次数
	openCaptcha := global.BIGO_CONFIG.Captcha.OpenCaptcha
	openCaptchaTimeOut := global.BIGO_CONFIG.Captcha.OpenCaptchaTimeOut
	// 从缓存中获取目标 ip，短时间内登录过几次
	value, ok := global.LocalCache.Get(clientIP)
	if !ok {
		// 第一次登录，加入缓存
		global.LocalCache.Set(clientIP, 1, time.Second*time.Duration(openCaptchaTimeOut))
	}
	// 返回是否需要校验验证码
	// 验证码每次都需要输入，或者登录超过了阈值则需要验证码
	oc = openCaptcha == 0 || openCaptcha < utils.InterfaceToInt(value)
	return
}
