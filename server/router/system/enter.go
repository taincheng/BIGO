package system

import v1 "server/api/v1"

type RouterGroup struct {
	UserRouter
}

var (
	baseApi = v1.ApiGroupApp.SystemApi.BaseApi
)
