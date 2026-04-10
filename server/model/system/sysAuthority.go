package system

import "time"

type SysAuthority struct {
	// 创建时间
	CreatedAt time.Time `gorm:""`
	// 更新时间
	UpdatedAt time.Time
	DeletedAt *time.Time `gorm:"index"`
	// 角色ID
	AuthorityId uint `json:"authorityId" binding:"required" gorm:"not null;unique;primary_key;comment:角色ID;size:90"`
	// 角色名
	AuthorityName string `json:"authorityName" binding:"required" gorm:"comment:角色名"`
	// 父角色ID
	ParentId *uint `json:"parentId" gorm:"comment:父角色ID"`
	// 角色之间多对多关系，gorm 自动创建 sys_data_authority_id 映射表
	DataAuthorityId []*SysAuthority `json:"dataAuthorityId" gorm:"many2many:sys_data_authority_id;"`
	// 子角色会在业务中动态读取，不写入表中
	Children      []SysAuthority `json:"children" gorm:"-"`
	SysBaseMenus  []*SysBaseMenu `json:"menus" gorm:"many2many:sys_authority_menus;"`
	Users         []*SysUser     `json:"-" gorm:"many2many:sys_user_authority;"`
	DefaultRouter string         `json:"defaultRouter" gorm:"comment:默认菜单;default:dashboard"` // 默认菜单(默认dashboard)
}

func (SysAuthority) TableName() string {
	return "sys_authorities"
}
