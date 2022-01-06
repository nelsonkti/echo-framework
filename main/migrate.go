package main

import (
	"echo-framework/lib/db"
	"echo-framework/logic/http/model"
	"fmt"
)

func main() {
	defer db.DisconnectMysql()
	db.InitMysql()

	err := db.Mysql("jz_ybs").AutoMigrate(&model.EmployeesBase{})
	fmt.Println(err)
}
