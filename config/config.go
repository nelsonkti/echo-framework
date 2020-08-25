package config

import "time"

// socket配置
var (
	ConnectTCPListenPort = "8000"
	LogicHTTPListenIP    = ":8080"
	AppSign              = "hello world 2"
	Env                  = "develop"
	AppDomain            = "xxxx"
	TimeZone, _          = time.LoadLocation("Asia/Shanghai")
)
