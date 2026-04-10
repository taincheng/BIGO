package system

import (
	"server/global"
	"server/model/system"
)

type LoginLogService struct{}

func (s *LoginLogService) CreateLoginLog(loginLog *system.SysLoginLog) (err error) {
	err = global.BIGO_DB.Create(loginLog).Error
	return
}
