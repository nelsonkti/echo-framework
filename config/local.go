package config

import (
	"github.com/tidwall/gjson"
	"io/ioutil"
	"echo-framework/lib/helper"
	"log"
)

var configStr string
var value string
var values []string

func init() {

	// 获取配置文件内容
	configStr = getContent()
	if configStr == "" {
		return
	}

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

	// 设置域名
	newDomain()
}

func getContent() string {
	content, err := ioutil.ReadFile("config.json")
	if err != nil {
		log.Println(err)
		return ""
	}

	return string(content)
}

// 设置环境
func env() {
	//环境
	value = gjson.Get(configStr, "env").String()

	if value != "" {
		Env = value
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
			write := i.Get("write").Array()
			if write != nil {
				for _, writeHost := range write {
					MysqlConfig[configName+MasterSuffix] = struct {
						Host string
						Port string
						User string
						Pwd  string
						Name string
					}{
						Host: writeHost.String(),
						Port: i.Get("port").String(),
						User: i.Get("user").String(),
						Pwd:  i.Get("pwd").String(),
						Name: i.Get("name").String(),
					}
				}

			}

			MysqlConfig[configName] = struct {
				Host string
				Port string
				User string
				Pwd  string
				Name string
			}{
				Host: i.Get("host").String(),
				Port: i.Get("port").String(),
				User: i.Get("user").String(),
				Pwd:  i.Get("pwd").String(),
				Name: i.Get("name").String(),
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
	for _, v := range gjson.Get(configStr, "nsq").Array() {
		values = append(values, v.String())
	}
	if len(values) > 0 {
		NSQConsumers = values
	}

	//nsq tcp hosts
	NSQServerHosts[helper.GetLocalIP()+"."+ConnectTCPListenPort] = struct{}{}
	for _, v := range gjson.Get(configStr, "nsq_server_hosts").Array() {
		NSQServerHosts[v.String()+"."+ConnectTCPListenPort] = struct{}{}
	}
}

func newDomain() {
	value = gjson.Get(configStr, "app_domain").String() //项目域名
	if value != "" {
		AppDomain = value
	}
}
