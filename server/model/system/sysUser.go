package system

import "github.com/google/uuid"

type SysUser struct {
	Model
	UUID        uuid.UUID    `json:"uuid" gorm:"index;comment:用户 UUID"`
	Username    string       `json:"username" gorm:"index;comment:用户名"`
	Password    string       `json:"-" gorm:"comment:密码"`
	AuthorityId uint         `json:"authorityId" gorm:"default:888;comment:用户角色ID"`
	Authority   SysAuthority `json:"authority" gorm:"foreignKey:AuthorityId;references:AuthorityId;comment:用户角色"`
}
