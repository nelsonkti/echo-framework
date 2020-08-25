package main

import (
	"echo-framework/config"
)


func main() {
	// 迁移伊把手
	//conn := db.ConnectMysql(getDbConfig("jz_ybs"), "jz_ybs")
	//defer db.DisconnectMysql()
	//conn.AutoMigrate(&model.EmployeesBase{})
}

/**
	获取Db配置
 */
func getDbConfig(dbName string) config.DBConfig {
	if dbConfig, ok := config.MysqlConfig[dbName]; ok {
		return dbConfig
	} else {
		defer println("该数据库【"+dbName+"】链接失败")
		panic("数据库断开连接")
	}
}
