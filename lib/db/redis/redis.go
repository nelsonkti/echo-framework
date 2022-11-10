package redis

import (
	"github.com/go-redis/redis"
	"github.com/nelsonkti/echo-framework/lib/logger"
	"sync"
)

//redis 服务
var redisDatabases sync.Map

type Option func(*options)

type DBFunc func() int
type NameFunc func() string

type options struct {
	db   DBFunc
	name NameFunc
}

func NewClient(address, password string, opts ...Option) *redis.Client {
	o := options{
		db:   defaultDB,
		name: defaultName,
	}
	for _, opt := range opts {
		opt(&o)
	}

	redisClient := redis.NewClient(
		&redis.Options{
			Addr:     address,
			Password: password,
			DB:       o.db(),
		},
	)

	_, err := redisClient.Ping().Result()
	if err != nil {
		panic(err)
	}

	redisDatabases.Store(o.name(), redisClient)

	return redisClient
}

func connect(name string) *redis.Client {
	value, ok := redisDatabases.Load(name)
	if ok {
		return value.(*redis.Client)
	}
	logger.Sugar.Info("failed to connect redis database:" + name)
	panic("failed to connect redis database:" + name)
}

func RedisDefault() *redis.Client {
	return connect("default")
}

func Disconnect() {
	redisDatabases.Range(func(key, value interface{}) bool {
		defer value.(*redis.Client).Close()
		return true
	})
	logger.Sugar.Info("disconnect redis")
}

func defaultDB() int {
	return 0
}

func WithDB(db int) Option {
	return func(o *options) {
		o.db = func() int {
			return db
		}
	}
}

func defaultName() string {
	return "default"
}

func WithName(name string) Option {
	return func(o *options) {
		o.name = func() string {
			return name
		}
	}
}
