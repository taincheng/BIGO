package initialize

import (
	"server/global"
)

// bizModel 迁移业务表
func bizModel() error {
	db := global.BIGO_DB
	err := db.AutoMigrate()
	if err != nil {
		return err
	}
	return nil
}
