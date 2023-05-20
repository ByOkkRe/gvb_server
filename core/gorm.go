package core

import (
	"gvb_server/global"
	"time"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

func Gorm() *gorm.DB {
	return MysqlConnect()
}

func MysqlConnect() *gorm.DB {
	if global.Config.Mysql.Host == "" {
		global.Log.Warn("未配置mysql，取消gorm连接")
		return nil
	}
	dsn := global.Config.Mysql.Dsn()
	var mysqlLogger logger.Interface
	if global.Config.System.Env == "dev" {
		mysqlLogger = logger.Default.LogMode(logger.Info)
	} else {
		mysqlLogger = logger.Default.LogMode(logger.Error)
	}

	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{
		Logger: mysqlLogger,
	})
	if err != nil {
		global.Log.Fatal("[%s] mysql连接失败", dsn)
		panic(err)
	}
	sqlDB, _ := db.DB()
	// SetMaxIdleConns 设置空闲连接池中的最大连接数。
	sqlDB.SetMaxIdleConns(10)
	// SetMaxOpenConns设置数据库的最大开放连接数	sqlDB.SetMaxOpenConns(100)
	sqlDB.SetMaxOpenConns(global.Config.Mysql.MaxOpenConns)
	// SetConnMaxLifetime设置一个连接可以被重复使用的最大时间。
	sqlDB.SetConnMaxLifetime(time.Hour)
	return db
}
