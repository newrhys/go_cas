package core

import (
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"net/url"
	"wave-admin/global"
)

func InitGorm() *gorm.DB {
	switch global.GnConfig.System.DbType {
	case "mysql":
		return GormMysql()
	default:
		return GormMysql()
	}
}

func GormMysql() *gorm.DB {
	m := global.GnConfig.Mysql
	if m.Dbname == "" {
		return nil
	}
	dsn := fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=%s&parseTime=True&loc=%s",
		m.Username,
		m.Password,
		m.Addr,
		m.Dbname,
		m.Charset,
		url.QueryEscape(m.Loc))

	mysqlConfig := mysql.Config{
		DSN:                       dsn,   // DSN data source name
		DefaultStringSize:         255,   // string 类型字段的默认长度
		DisableDatetimePrecision:  false,  // 禁用 datetime 精度，MySQL 5.6 之前的数据库不支持
		DontSupportRenameIndex:    true,  // 重命名索引时采用删除并新建的方式，MySQL 5.7 之前的数据库和 MariaDB 不支持重命名索引
		DontSupportRenameColumn:   true,  // 用 `change` 重命名列，MySQL 8 之前的数据库和 MariaDB 不支持重命名列
		SkipInitializeWithVersion: false, // 根据版本自动配置
	}

	db, err := gorm.Open(mysql.New(mysqlConfig), gormConfig(m.LogMode))
	if err != nil {
		panic("failed to connect database, err: " + err.Error())
		return nil
	}

	sqlDB, _ := db.DB()
	sqlDB.SetMaxIdleConns(m.MaxIdleConns)
	sqlDB.SetMaxOpenConns(m.MaxOpenConns)

	global.GnLog.Info("数据库连接成功")

	return db
}

func gormConfig(mod bool) *gorm.Config {
	var config = &gorm.Config{DisableForeignKeyConstraintWhenMigrating: true}
	switch global.GnConfig.Mysql.LogZap {
	case "silent", "Silent":
		config.Logger = logger.Default.LogMode(logger.Silent)
	case "error", "Error":
		config.Logger = logger.Default.LogMode(logger.Error)
	case "warn", "Warn":
		config.Logger = logger.Default.LogMode(logger.Warn)
	case "info", "Info":
		config.Logger = logger.Default.LogMode(logger.Info)
	case "zap", "Zap":
		config.Logger = logger.Default.LogMode(logger.Info)
	default:
		if mod {
			config.Logger = logger.Default.LogMode(logger.Info)
			break
		}
		config.Logger = logger.Default.LogMode(logger.Silent)
	}
	return config
}