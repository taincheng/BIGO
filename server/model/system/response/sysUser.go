package response

import "server/model/system"

type SysUserResponse struct {
	User system.SysUser `json:"user"`
}
