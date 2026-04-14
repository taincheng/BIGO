package system

import v1 "server/api/v1"

type RouterGroup struct {
	UserRouter
	BaseRouter
	AuthorityRouter
}

var (
	baseApi      = v1.ApiGroupApp.SystemApi.BaseApi
	authorityApi = v1.ApiGroupApp.SystemApi.AuthorityApi
)
