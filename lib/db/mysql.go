package db

import (
	"fmt"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	config2 "echo-framework/config"
	"echo-framework/lib/logger"
	"sync"
	"time"
)

var mysqlDatabases sync.Map

const defaultConfig = "?parseTime=true&charset=utf8mb4&loc=Asia%2FShanghai"

func InitMysql() {
	for key, Config := range config2.MysqlConfig {
		ConnectMysql(Config, key)
	}
}

func ConnectMysql(Config config2.DBConfig, name string) *gorm.DB {
	sql := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s%s", Config.User, Config.Pwd, Config.Host, Config.Port, Config.Name, defaultConfig)

	db, _ := gorm.Open("mysql", sql)

	err := db.DB().Ping()
	if err != nil {
		panic("failed to connect mysql:" + name)
	}

	db.DB().SetMaxIdleConns(1024)
	db.DB().SetMaxOpenConns(1024)
	db.DB().SetConnMaxLifetime(time.Minute * 10) //连接超时10分钟，数据库的wait_timeout最好设置为11分钟

	if config2.Env == "local" {
		db.LogMode(true)
	}

	mysqlDatabases.Store(name, db)

	// 启用Logger，显示详细日志
	return db
}

func Mysql(name string) *gorm.DB {
	db, _ := mysqlDatabases.Load(name)
	return db.(*gorm.DB)
}

//
func ModelMaster(name string) *gorm.DB {
	db := Mysql(name + config2.MasterSuffix)
	if db.DB().Ping() != nil {
		return Mysql(name)
	}
	return db
}

func DisconnectMysql() {
	mysqlDatabases.Range(func(key, value interface{}) bool {
		defer value.(*gorm.DB).Close()
		return true
	})
	logger.Sugar.Info("disconnect mysql")
}
