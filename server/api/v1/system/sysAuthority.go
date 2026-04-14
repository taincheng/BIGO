package system

import (
	"server/global"
	"server/model/common/response"
	"server/model/system"
	systemRes "server/model/system/response"
	"server/utils"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

type AuthorityApi struct{}

// CreateAuthority
// @Tags      Authority
// @Summary   创建角色
// @Security  ApiKeyAuth
// @accept    application/json
// @Produce   application/json
// @Param     data  body      system.SysAuthority                                                true  "权限id, 权限名, 父角色id"
// @Success   200   {object}  response.Response{data=systemRes.SysAuthorityResponse,msg=string}  "创建角色,返回包括系统角色详情"
// @Router    /authority/createAuthority [post]
func (a *AuthorityApi) CreateAuthority(c *gin.Context) {
	var authority system.SysAuthority
	if err := c.ShouldBindJSON(&authority); err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}

	if (authority.ParentId == nil || *authority.ParentId == 0) && global.BIGO_CONFIG.System.UseStrictAuth {
		authority.ParentId = utils.Ptr(utils.GetUserAuthorityId(c))
	}

	if err := authorityService.CreateAuthority(&authority); err != nil {
		global.BIGO_LOG.Error("创建角色失败！", zap.Error(err))
		response.FailWithMessage("创建角色失败:"+err.Error(), c)
		return
	}

	err := casbinService.FreshCasbin()
	if err != nil {
		global.BIGO_LOG.Error("创建成功，权限刷新失败", zap.Error(err))
		response.FailWithMessage("创建成功，权限刷新失败: "+err.Error(), c)
		return
	}
	response.OkWithDetailed(systemRes.SysAuthorityResponse{Authority: authority}, "创建成功", c)
}
