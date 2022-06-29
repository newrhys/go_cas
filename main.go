package main

import (
	"wave-admin/core"
	"wave-admin/global"
	"wave-admin/initialize"
)

// @title gin框架后台管理demo
// @version 1.0
// @description gin框架后台管理
// @host 127.0.0.1:9111

func main() {
	global.GnVp = core.InitViper()   // 初始化Viper
	global.GnLog = core.InitLogger() // 初始化zap日志库
	global.GnTrans = core.InitTrans()
	global.GnDb = core.InitGorm() // 初始化数据库
	if global.GnDb != nil {
		// 初始化表
		//initialize.MysqlTables(global.GnDb)
		global.GnCasbin, _ = initialize.Casbin()
		// 程序结束前关闭数据库链接
		db, _ := global.GnDb.DB()
		defer db.Close()
	}

	core.RunServer()
}
