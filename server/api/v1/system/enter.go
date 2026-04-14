package system

import "server/service"

type ApiGroup struct {
	BaseApi
	AuthorityApi
}

var (
	userService      = service.ServiceGroupApp.SystemServiceGroup.UserService
	LoginLogService  = service.ServiceGroupApp.SystemServiceGroup.LoginLogService
	authorityService = service.ServiceGroupApp.SystemServiceGroup.AuthorityService
	casbinService    = service.ServiceGroupApp.SystemServiceGroup.CasbinService
)
