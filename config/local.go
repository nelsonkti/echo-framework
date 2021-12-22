package config

import (
	"echo-framework/lib/localtion"
	"echo-framework/lib/logger"
	"github.com/tidwall/gjson"
	"io/ioutil"
	"os"
	"strings"
)

var configStr string
var value string
var values []string

func init() {

	setAppPath()

	// 获取配置文件内容
	getContent()

	// 设置环境
	env()

	// 设置mysql
	newMysql()

	// 设置redis
	newRedis()

	//memcache 配置
	newMemcache()

	//消息队列
	newNsq()

	// 日志
	newLogger()
}

func setAppPath() {

	path, err := os.Getwd()

	if err != nil {
		panic(err)
	}


	if strings.Contains(path, AppName) {
		AppPath = path[0 : strings.Index(path, AppName)+len(AppName)]
	} else {
		AppPath = path
	}

}

func getContent() {

	content, err := ioutil.ReadFile(AppPath + "/config.json")

	if err != nil {
		//logger.Sugar.Error(err)
		content, err = ioutil.ReadFile("/wwwroot/config.json")
		if err != nil {
			//logger.Sugar.Error(err)
			panic(err)
		}
	}

	configStr = string(content)

}

// 设置环境
func env() {
	//环境
	value = gjson.Get(configStr, "env").String()

	if value != "" {
		Env = value
	}

	value = gjson.Get(configStr, "port").String()

	if value != "" {
		LogicHTTPListenIP = ":" + value
	}
}

//设置mysql
func newMysql() {
	//mysql 配置
	result := gjson.Get(configStr, "mysql")
	for _, name := range result.Array() {
		for _, i := range name.Array() {
			configName := i.Get("conn_name").String()

			// 检查是否存在主从
			read := i.Get("read").Array()

			var rHost []string

			if read != nil {
				for _, writeHost := range read {
					rHost = append(rHost, writeHost.String())
				}
			}

			MysqlConfig[configName] = struct {
				Host  string
				Port  string
				User  string
				Pwd   string
				Name  string
				RHost []string
			}{
				Host:  i.Get("host").String(),
				Port:  i.Get("port").String(),
				User:  i.Get("user").String(),
				Pwd:   i.Get("pwd").String(),
				Name:  i.Get("name").String(),
				RHost: rHost,
			}
		}
	}
}

// 设置redis
func newRedis() {
	value = gjson.Get(configStr, "redis_address").String()

	if value != "" {
		RedisIP = value
	}
	value = gjson.Get(configStr, "redis_password").String()
	if value != "" {
		RedisPassword = value
	} else if Env == "local" {
		RedisPassword = value
	}
}

// 设置 memcache
func newMemcache() {
	values = make([]string, 0)
	for _, v := range gjson.Get(configStr, "memcache").Array() {
		values = append(values, v.String())
	}
	if len(values) > 0 {
		Memcache = values
	}
}

//
func newNsq() {
	values = make([]string, 0)
	for _, v := range gjson.Get(configStr, "nsq_consumers").Array() {
		values = append(values, v.String())
	}
	if len(values) > 0 {
		NSQConsumers = values
	}

	NSQServerHost := gjson.Get(configStr, "nsq_server_hosts").String()

	if NSQServerHost != "" {
		NSQIP = NSQServerHost
	}

	//nsq tcp hosts
	NSQServerHosts[localtion.GetLocalIP()+"."+ConnectTCPListenPort] = struct{}{}
	for _, v := range gjson.Get(configStr, "nsq_server_hosts").Array() {
		NSQServerHosts[v.String()+"."+ConnectTCPListenPort] = struct{}{}
	}

}


// 初始化日志文件
func newLogger() {
	logger.InitLogger(AppName)
}


