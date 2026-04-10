package system

type SysMenu struct {
	SysBaseMenu
	MenuId      uint                   `json:"menuId"`
	AuthorityId uint                   `json:"-"`
	Children    []SysMenu              `json:"children"`
	Parameters  []SysBaseMenuParameter `json:"parameters"`
	Btns        map[string]uint        `json:"btns"`
}

type SysAuthorityMenu struct {
	// 菜单ID
	MenuId string `json:"menuId"`
	// 角色ID
	AuthorityId string `json:"-"`
}
