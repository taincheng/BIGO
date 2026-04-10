package request

import "server/model/system"

// DefaultMenu 默认菜单
func DefaultMenu() []system.SysBaseMenu {
	return []system.SysBaseMenu{
		{
			Model: system.Model{
				ID: 1,
			},
			ParentId:  0,
			Path:      "dashboard",
			Name:      "dashboard",
			Component: "view/dashboard/index.vue",
			Sort:      1,
			Meta: system.Meta{
				Title: "仪表盘",
				Icon:  "setting",
			},
		},
	}
}
