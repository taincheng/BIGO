package system

import (
	"server/global"
	"server/model/common/response"
	"server/model/system"
	systemReq "server/model/system/request"
	systemRes "server/model/system/response"
	"server/utils"
	"time"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

type BaseApi struct{}

// Register
// @Tags sysUser
// @Summary 用户注册
// @Produce application/json
// @Param data body systemReq.Register true "用户名, 密码"
// @Success 200 {object} response.Response{data=systemRes.SysUserResponse,msg=string} "注册用户返回信息"
// @Router /user/admin_register [post]
func (b *BaseApi) Register(c *gin.Context) {
	var r systemReq.Register
	err := c.ShouldBindJSON(&r)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	var authorityIds []*system.SysAuthority
	for _, v := range r.AuthorityIds {
		authorityIds = append(authorityIds, &system.SysAuthority{
			AuthorityId: v,
		})
	}
	user := system.SysUser{
		Username:    r.Username,
		Password:    r.Password,
		AuthorityId: r.AuthorityId,
		Authorities: authorityIds,
		Phone:       r.Phone,
		Email:       r.Email,
		Enable:      r.Enable,
	}
	err = userService.Register(&user)
	if err != nil {
		global.BIGO_LOG.Error("注册失败！", zap.Error(err))
		response.FailWithDetailed(systemRes.SysUserResponse{User: user}, "注册失败", c)
		return
	}
	response.OkWithDetailed(systemRes.SysUserResponse{User: user}, "注册成功", c)
}

// Login
// @Tags SysUser
// @Summary 用户登录
// @Produce  application/json
// @Param data body systemReq.LoginReq true "用户名, 密码"
func (b *BaseApi) Login(c *gin.Context) {
	var login systemReq.Login
	if err := c.ShouldBindJSON(login); err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	clientIP := c.ClientIP()
	// 判断验证码是否开启
	openCaptcha := global.BIGO_CONFIG.Captcha.OpenCaptcha
	openCaptchaTimeOut := global.BIGO_CONFIG.Captcha.OpenCaptchaTimeOut
	// 从缓存中获取目标 ip 获取过几次验证码
	value, ok := global.LocalCache.Get(clientIP)
	if !ok {
		// 第一次请求验证码，加入缓存
		global.LocalCache.Set(clientIP, 1, time.Second*time.Duration(openCaptchaTimeOut))
	}
	// 获取验证码的次数是否达到防爆次数
	var oc bool = openCaptcha == 0 || openCaptcha < utils.InterfaceToInt(value)
	// TODO
	//if oc && (login.Captcha == "" || login.CaptchaId == "" || )
}
