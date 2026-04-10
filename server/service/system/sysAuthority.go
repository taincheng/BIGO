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
		if err := global.BIGO_DB.Create(auth).Error; err != nil {
			return err
		}
		auth.SysBaseMenus = utils.ToPtrSlice(request.DefaultMenu())

		if err := tx.Model(auth).Association("SysBaseMenus").Replace(auth.SysBaseMenus); err != nil {
			return err
		}
		casbinInfos := request.DefaultCasbin()
		authorityId := strconv.Itoa(int(auth.AuthorityId))
		rules := make([][]string, 0, len(casbinInfos))
		for _, v := range casbinInfos {
			rules = append(rules, []string{authorityId, v.Path, v.Method})
		}
		// TODO
		return nil
	})
}
