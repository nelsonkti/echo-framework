package main

import (
	"fmt"
	"github.com/nelsonkti/echo-framework/config"
	"github.com/nelsonkti/echo-framework/lib/db/mysql"
	"github.com/nelsonkti/echo-framework/lib/db/mysql/db"
	"github.com/nelsonkti/echo-framework/lib/logger"
	"github.com/nelsonkti/echo-framework/logic/http/model"
)

func main() {
	// 日志
	logger.New(logger.Base(config.AppConf.App.Env, config.AppConf.App.Path.LogPath))

	defer mysql.Disconnect()
	mysql.Connect()

	err := db.Connect("db").AutoMigrate(&model.EmployeesBase{})
	fmt.Println(err)
}
