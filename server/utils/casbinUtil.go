package utils

import (
	"server/global"
	"sync"

	"github.com/casbin/casbin/v3"
	"github.com/casbin/casbin/v3/model"
	gormadapter "github.com/casbin/gorm-adapter/v3"
)

var (
	syncedCachedEnforcer *casbin.SyncedCachedEnforcer
	once                 sync.Once
)

// GetCasbin 获取 casbin 实例
func GetCasbin() *casbin.SyncedCachedEnforcer {
	once.Do(func() {
		db, err := gormadapter.NewAdapterByDB(global.BIGO_DB)
		if err != nil {
			global.BIGO_LOG.Error(err.Error())
			return
		}
		text := `
		[request_definition]
		r = sub, obj, act
		
		[policy_definition]
		p = sub, obj, act
		
		[role_definition]
		g = _, _
		
		[policy_effect]
		e = some(where (p.eft == allow))
		
		[matchers]
		m = r.sub == p.sub && keyMatch2(r.obj,p.obj) && r.act == p.act
		`
		fromString, err := model.NewModelFromString(text)
		if err != nil {
			global.BIGO_LOG.Error("加载 casbin 模型失败")
			return
		}
		syncedCachedEnforcer, _ = casbin.NewSyncedCachedEnforcer(fromString, db)
		syncedCachedEnforcer.SetExpireTime(60 * 60)
		err = syncedCachedEnforcer.LoadPolicy()
		if err != nil {
			global.BIGO_LOG.Error("加载 casbin 策略失败")
			return
		}
	})
	return syncedCachedEnforcer
}
