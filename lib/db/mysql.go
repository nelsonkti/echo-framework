package db

import (
	"database/sql"
	"echo-framework/config"
	my_logger "echo-framework/lib/logger"
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"gorm.io/plugin/dbresolver"
	_ "gorm.io/plugin/dbresolver"
	"sync"
	"time"
)

var mysqlDatabases sync.Map

const defaultConfig = "?parseTime=true&charset=utf8mb4&loc=Asia%2FShanghai"

func InitMysql() {
	for key, Config := range config.MysqlConfig {
		ConnectMysql(Config, key)
	}
}

func ConnectMysql(Config config.DBConfig, name string) *gorm.DB {
	var dialect []gorm.Dialector

	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s%s", Config.User, Config.Pwd, Config.Host, Config.Port, Config.Name, defaultConfig)

	logLevel := logger.Silent

	if config.Env == "local" {
		logLevel = logger.Info
	}

	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{
		Logger: logger.Default.LogMode(logLevel),
	})

	if err != nil {
		my_logger.Sugar.Info(err)
		panic(err)
	}

	if Config.RHost != nil {

		for _, v := range Config.RHost {
			sqlRead := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s%s", Config.User, Config.Pwd, v, Config.Port, Config.Name, defaultConfig)
			dialect = append(dialect, mysql.Open(sqlRead))
		}

		_ = db.Use(dbresolver.Register(dbresolver.Config{
			Sources:  []gorm.Dialector{mysql.Open(dsn)},
			Replicas: dialect,
			//  负载均衡策略
			Policy: dbresolver.RandomPolicy{},
		}))

	}

	sqlDB, err := db.DB()

	if err != nil {
		my_logger.Sugar.Info("failed to connect mysql:" + name)
		panic("failed to connect mysql:" + name)
	}

	ping(sqlDB, name)

	sqlDB = setDefault(sqlDB)

	mysqlDatabases.Store(name, db)

	// 启用Logger，显示详细日志
	return db
}

func Mysql(name string) *gorm.DB {
	db, _ := mysqlDatabases.Load(name)
	return db.(*gorm.DB)
}

func DisconnectMysql() {

	mysqlDatabases.Range(func(key, value interface{}) bool {
		db, _ := value.(*gorm.DB)
		sqlDB, _ := db.DB()
		sqlDB.Close()
		return true
	})

	my_logger.Sugar.Info("disconnect mysql")
}

func ping(sqlDB *sql.DB, name string) {
	err := sqlDB.Ping()
	if err != nil {
		my_logger.Sugar.Info("failed to connect mysql:" + name)
		panic("failed to connect mysql:" + name)
	}
}

func setDefault(sqlDB *sql.DB) *sql.DB {

	sqlDB.SetMaxIdleConns(1024)

	sqlDB.SetMaxOpenConns(1024)

	sqlDB.SetConnMaxLifetime(time.Minute * 10)

	return sqlDB
}
