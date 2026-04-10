package system

import (
	"errors"
	"server/global"
	"server/model/system"

	"gorm.io/gorm"
)

type MenuService struct{}

var MenuServiceApp = new(MenuService)

func (m *MenuService) UserAuthorityDefaultRouter(user *system.SysUser) {
	var menuIds []string
	err := global.BIGO_DB.Model(&system.SysAuthorityMenu{}).
		Where("sys_authority_authority_id = ?", user.AuthorityId).
		Pluck("sys_base_menu_id", &menuIds).Error
	if err != nil {
		return
	}
	var baseMenu []system.SysBaseMenu
	err = global.BIGO_DB.First(&baseMenu, "name = ? and id in (?)", user.Authority.DefaultRouter, menuIds).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		user.Authority.DefaultRouter = "404"
	}
}
