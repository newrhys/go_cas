package initialize

import (
	"github.com/casbin/casbin/v2"
	gormadapter "github.com/casbin/gorm-adapter/v3"
	"go.uber.org/zap"
	"log"
	"wave-admin/global"
)

func Casbin() (syncedEnforcer *casbin.SyncedEnforcer, err error) {
	a, _ := gormadapter.NewAdapterByDB(global.GnDb)
	syncedEnforcer, err = casbin.NewSyncedEnforcer(global.GnConfig.Casbin.ModelPath, a)
	if err != nil {
		global.GnLog.Error("Casbin加载失败!", zap.Error(err))
		log.Println(err)
		return nil, err
	} else {
		global.GnLog.Error("Casbin加载成功!")
	}
	_ = syncedEnforcer.LoadPolicy()

	return syncedEnforcer, err
}
