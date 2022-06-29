package initialize

import (
	"go.uber.org/zap"
	"gorm.io/gorm"
	"os"
	"wave-admin/global"
	"wave-admin/model/system"
)

func MysqlTables(db *gorm.DB) {
	//err := db.AutoMigrate(system.SysUser{}, system.JwtBlacklist{}, system.SysMenu{}, system.SysRole{}, system.SysRoleMenu{}, system.SysUserRole{}, system.SysApi{}, system.SysRoleApi{}, system.SysIcon{}, cms.CmsCategory{}, cms.CmsTag{}, cms.CmsArticle{}, cms.CmsArticleContent{}, cms.CmsArticleTag{}, adapter.CasbinRule{})
	err := db.AutoMigrate(system.SysRecord{})
	if err != nil {
		global.GnLog.Error("register table failed", zap.Any("err", err))
		os.Exit(0)
	}
	global.GnLog.Info("register table success")
}