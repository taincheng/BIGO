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
// @Param data body systemReq.Login true "用户名, 密码, 验证码, 验证码 id"
// @Success 200 {object} response.Response{data=systemRes.LoginResponse,msg=string} "包含用户信息，token, 登录过期时间"
// @Router /user/login [post]
func (b *BaseApi) Login(c *gin.Context) {
	var login systemReq.Login
	if err := c.ShouldBindJSON(&login); err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	// 判断是否需要校验验证码
	oc := b.judgeUseCaptcha(c)
	// 验证码校验失败
	if oc && (login.Captcha == "" || login.CaptchaId == "" || !store.Verify(login.CaptchaId, login.Captcha, true)) {
		// 获取验证码的次数+1
		global.LocalCache.Increment(c.ClientIP(), 1)
		response.FailWithMessage("验证码错误", c)
		b.loginLog(c, false, "验证码错误")
	}
	userEntry := &system.SysUser{
		Username: login.Username,
		Password: login.Password,
	}
	userInfo, err := userService.Login(userEntry)
	if err != nil {
		global.BIGO_LOG.Error("登录失败！用户名不存在或者密码错误", zap.Error(err))
		global.LocalCache.Increment(c.ClientIP(), 1)
		response.FailWithMessage("登录失败！用户名不存在或者密码错误", c)
		b.loginLog(c, false, "用户名不存在或者密码错误")
		return
	}
	if userInfo.Enable != 1 {
		global.BIGO_LOG.Error("登录失败！用户被禁用", zap.Error(err))
		global.LocalCache.Increment(c.ClientIP(), 1)
		response.FailWithMessage("登录失败！用户被禁用", c)
		b.loginLog(c, false, "用户被禁用")
		return
	}
	b.tokenNext(c, userInfo)
}

// tokenNext 签发 Token
func (b *BaseApi) tokenNext(c *gin.Context, user *system.SysUser) {
	jwt := utils.NewJWT()
	claims := jwt.CreateClaims(systemReq.BaseClaims{
		UUID:        user.UUID,
		ID:          user.ID,
		Username:    user.Username,
		AuthorityId: user.AuthorityId,
	})
	token, err := jwt.CreateToken(claims)
	if err != nil {
		global.BIGO_LOG.Error("获取token失败！", zap.Error(err))
		response.FailWithMessage("获取token失败", c)
		return
	}
	// 记录登录成功日志
	b.loginLog(c, true, "登录成功")
	// token 写入到 cookie 中
	utils.SetToken(c, token, int(claims.RegisteredClaims.ExpiresAt.Unix()-time.Now().Unix()))
	response.OkWithDetailed(systemRes.LoginResponse{
		User:      user,
		Token:     token,
		ExpiresAt: claims.RegisteredClaims.ExpiresAt.Unix() * 1000,
	}, "登录成功", c)
	return
}

func (b *BaseApi) loginLog(c *gin.Context, status bool, errorMessage string) {
	err := LoginLogService.CreateLoginLog(&system.SysLoginLog{
		Username:     c.PostForm("username"),
		Ip:           c.ClientIP(),
		Agent:        c.Request.UserAgent(),
		Status:       status,
		ErrorMessage: errorMessage,
	})
	if err != nil {
		global.BIGO_LOG.Error("登录日志创建失败！", zap.Error(err))
	}
}
