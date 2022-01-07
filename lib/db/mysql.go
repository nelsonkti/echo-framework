package db

import (
	"database/sql"
	"echo-framework/config"
	pb "echo-framework/config/pb"
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

var Databases sync.Map

const defaultConfig = "?parseTime=true&charset=utf8mb4&loc=Asia%2FShanghai"

func InitMysql() {
	for _, conf := range config.AppConf.Data.Connection.Database {
		newDatabase().connect(conf)
	}
}

func Mysql(name string) *gorm.DB {
	db, _ := Databases.Load(name)
	return db.(*gorm.DB)
}

func DisconnectMysql() {

	Databases.Range(func(key, value interface{}) bool {
		db, _ := value.(*gorm.DB)
		sqlDB, _ := db.DB()
		_ = sqlDB.Close()
		return true
	})

	my_logger.Sugar.Info("disconnect mysql")
}

func newDatabase() *Db {
	db := &Db{}

	db.dbConfig = db.defaultConfig
	db.connSlave = db.connectSlave
	db.ping = db.defaultPing

	return db
}

type nilFunc func()
type connSlaveFunc func(s string)

type Db struct {
	db        *gorm.DB
	sqlDb     *sql.DB
	database  *pb.Data_Database
	dbConfig  nilFunc
	connSlave connSlaveFunc
	ping      nilFunc
}

func (d *Db) connect(database *pb.Data_Database) {
	d.database = database

	dsn := tcpSprint(database, database.Host)

	var err error
	d.db, err = gorm.Open(mysql.Open(dsn), d.config())

	if err != nil {
		my_logger.Sugar.Info(err)
		panic(err)
	}

	d.connSlave(dsn)

	d.DB()
	d.ping()

	d.dbConfig()

	Databases.Store(database.Database, d.db)
}

func (d *Db) config() *gorm.Config {
	logLevel := logger.Silent

	if config.AppConf.App.Env == "local" {
		logLevel = logger.Info
	}

	return &gorm.Config{
		Logger: logger.Default.LogMode(logLevel),
	}
}

func (d *Db) connectSlave(dsn string) {
	if d.database.Read == nil {
		return
	}

	var dialect []gorm.Dialector
	for _, v := range d.database.Read {
		dialect = append(dialect, mysql.Open(tcpSprint(d.database, v)))
	}

	err := d.db.Use(dbresolver.Register(dbresolver.Config{
		Sources:  []gorm.Dialector{mysql.Open(dsn)},
		Replicas: dialect,
		//  负载均衡策略
		Policy: dbresolver.RandomPolicy{},
	}))

	if err != nil {
		panic(err)
	}

}

func (d *Db) DB() (Db *Db) {
	sqlDb, err := d.db.DB()
	if err != nil {
		my_logger.Sugar.Info("failed to connect mysql:" + d.database.Database)
		panic("failed to connect mysql:" + d.database.Database)
	}
	d.sqlDb = sqlDb

	return
}

func (d *Db) defaultPing() {
	err := d.sqlDb.Ping()
	if err != nil {
		my_logger.Sugar.Info("failed to connect mysql:" + d.database.Database)
		panic("failed to connect mysql:" + d.database.Database)
	}
}

func (d *Db) defaultConfig() {

	d.sqlDb.SetMaxIdleConns(1024)

	d.sqlDb.SetMaxOpenConns(1024)

	d.sqlDb.SetConnMaxLifetime(time.Minute * 10)

}

func tcpSprint(conf *pb.Data_Database, network string) string {
	return fmt.Sprintf("%s:%s@tcp(%s:%d)/%s%s", conf.Username, conf.Password, network, conf.Port, conf.Database, defaultConfig)
}
