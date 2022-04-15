package main

import (
	"fmt"
	"github.com/nelsonkti/echo-framework/config"
	"github.com/nelsonkti/echo-framework/lib/db"
	"github.com/nelsonkti/echo-framework/lib/logger"
	"github.com/nelsonkti/echo-framework/logic/http/model"
)

func main() {
	// 日志
	logger.New(logger.Base(config.AppConf.App.Env, config.AppConf.App.Path.LogPath))

	defer db.DisconnectMysql()
	db.InitMysql()

	err := db.Mysql("db").AutoMigrate(&model.EmployeesBase{})
	fmt.Println(err)
}
