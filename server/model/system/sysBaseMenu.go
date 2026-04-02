package system

type SysBaseMenu struct {
	Model
	MenuLevel uint   `json:"-" gorm:"comment:菜单层级"`
	ParentId  uint   `json:"parentId" gorm:"comment:父菜单ID"`
	Path      string `json:"path" gorm:"comment:路由path"`
	Name      string `json:"name" gorm:"comment:路由name"`
	Hidden    bool   `json:"hidden" gorm:"comment:是否在列表隐藏"`
	Component string `json:"component" gorm:"comment:对应前端文件路径"`
	Sort      int    `json:"sort" gorm:"comment:排序标记"`
	// 把 Meta 中的字段嵌入到该表中，不创建新的关联表
	Meta           `json:"meta" gorm:"embedded;comment:附加属性"`
	SysAuthorities []*SysAuthority        `json:"authorities" gorm:"many2many:sys_authority_menus;"`
	Children       []SysBaseMenu          `json:"children" gorm:"-"`
	Parameters     []SysBaseMenuParameter `json:"parameters"`
	MenuBtn        []SysBaseMenuBtn       `json:"menuBtn"`
}

type Meta struct {
	ActiveName string `json:"activeName" gorm:"comment:高亮菜单"`
	KeepAlive  bool   `json:"keepAlive" gorm:"comment:是否缓存"`
	Title      string `json:"title" gorm:"comment:菜单名"`
	Icon       string `json:"icon" gorm:"comment:菜单图标"`
	CloseTab   bool   `json:"closeTab" gorm:"comment:自动关闭tab"`
}

type SysBaseMenuParameter struct {
	Model
	SysBaseMenuID uint
	Type          string `json:"type" gorm:"comment:地址栏携带参数为params还是query"`
	Key           string `json:"key" gorm:"comment:地址栏携带参数的key"`
	Value         string `json:"value" gorm:"comment:地址栏携带参数的值"`
}

func (SysBaseMenu) TableName() string {
	return "sys_base_menus"
}
