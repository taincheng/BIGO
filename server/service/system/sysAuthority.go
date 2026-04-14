package system

import (
	"server/global"
	"server/model/system"
	"server/model/system/request"
	"server/utils"
	"strconv"

	"gorm.io/gorm"
)

type AuthorityService struct{}

var AuthorityServiceApp = new(AuthorityService)

// CreateAuthority 创建角色
func (a *AuthorityService) CreateAuthority(auth *system.SysAuthority) error {
	err := global.BIGO_DB.Transaction(func(tx *gorm.DB) error {
		// 创建角色
		if err := global.BIGO_DB.Create(auth).Error; err != nil {
			return err
		}
		auth.SysBaseMenus = utils.ToPtrSlice(request.DefaultMenu())
		// 清理关联表中的关联关系，将默认的关联数据替代进关联表
		if err := tx.Model(auth).Association("SysBaseMenus").Replace(auth.SysBaseMenus); err != nil {
			return err
		}
		// 设置默认 Casbin 权限规则
		casbinInfos := request.DefaultCasbin()
		authorityId := strconv.Itoa(int(auth.AuthorityId))
		// 将角色ID与每个权限路径、方法组合成 Casbin 规则
		rules := make([][]string, 0, len(casbinInfos))
		for _, v := range casbinInfos {
			rules = append(rules, []string{authorityId, v.Path, v.Method})
		}
		return CasbinServiceApp.AddPolicies(tx, rules)
	})
	return err
}
