package v1

import "server/api/v1/system"

var ApiGroupApp = new(ApiGroup)

type ApiGroup struct {
	SystemApi system.ApiGroup
}
