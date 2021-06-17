package config

const masterSuffixName = "_master"


type DBConfig struct{
	Host string
	Port string
	User string
	Pwd  string
	Name string
	RHost []string
}

// 数据库和服务配置
var (
	MysqlConfig = make(map[string]DBConfig)

	MasterSuffix = masterSuffixName

	Memcache = []string{"0.0.0.0:11211"}

	RedisIP       = "127.0.0.1:6379"
	RedisPassword = ""
)

var (
	NSQIP        = "0.0.0.0:4150"
	NSQConsumers = []string{
		"0.0.0.0:4161",
	}

	NSQServerHosts = make(map[string]struct{})
)
