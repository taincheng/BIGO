package system

import (
	"server/utils"

	gormadapter "github.com/casbin/gorm-adapter/v3"
	"gorm.io/gorm"
)

type CasbinService struct{}

var CasbinServiceApp = new(CasbinService)

// AddPolicies 添加匹配的权限
func (c *CasbinService) AddPolicies(db *gorm.DB, rules [][]string) error {
	var casbinRules []gormadapter.CasbinRule
	for i := range rules {
		casbinRules = append(casbinRules, gormadapter.CasbinRule{
			Ptype: "p",
			V0:    rules[i][0],
			V1:    rules[i][1],
			V2:    rules[i][2],
		})
	}
	return db.Create(&casbinRules).Error
}

func (c *CasbinService) FreshCasbin() error {
	casbin := utils.GetCasbin()
	err := casbin.LoadPolicy()
	return err
}
