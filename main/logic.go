package main

import (
	"echo-framework/config"
	"echo-framework/cron"
	"echo-framework/lib/db"
	"echo-framework/lib/helper"
	"echo-framework/lib/localtion"
	"echo-framework/lib/logger"
	"echo-framework/logic/nsq/consumer"
	"echo-framework/logic/nsq/producer"
	"echo-framework/routes"
	"echo-framework/util/xnsq"
	"echo-framework/util/xnsq/service/registry"
	"fmt"
	"github.com/judwhite/go-svc"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"os"
	"path/filepath"
	"sync"
	"syscall"
)

var Echo = echo.New()

type logicProgram struct {
	once sync.Once
}

func main() {
	p := &logicProgram{}
	if err := svc.Run(p, syscall.SIGINT, syscall.SIGTERM, syscall.SIGKILL); err != nil {
		logger.Sugar.Fatal(err)
	}
}

// svc 服务运行框架 程序启动时执行Init+Start, 服务终止时执行Stop
func (p *logicProgram) Init(env svc.Environment) error {
	if env.IsWindowsService() {
		dir := filepath.Dir(os.Args[0])
		return os.Chdir(dir)
	}
	return nil
}

func (p *logicProgram) Start() error {
	//连接mysql

	db.InitMysql()

	//连接 memcache
	db.ConnectMemcache(config.AppConf.Data.Memcache.Host)

	//连接redis
	db.ConnectRedis(config.AppConf.Data.Redis.Addr, config.AppConf.Data.Redis.Password, 0, "default")

	go func() {
		defer helper.RecoverPanic()

		server := xnsq.NewNsqServer(registry.Options{
			NsqAddress:   config.AppConf.Mq.Nsq.Host,
			NSQConsumers: config.AppConf.Mq.Nsq.Consumer,
			Env:          config.AppConf.App.Env,
			LocalAddress: localtion.GetLocalIP(),
		})

		server.Run(consumer.LogicConsumerHandler(server.Opt))

	}()

	//启动定时任务
	if config.AppConf.App.Env != "local" {
		cron.RegisterCrons(config.AppConf.Data.Redis.Addr, config.AppConf.Data.Redis.Password)
	}

	// 启动app
	go func() {
		newApp()
	}()

	return nil
}

/**
启动app
*/
func newApp() {
	e := Echo

	e.Binder = new(helper.Binder)

	e.Use(middleware.Recover())
	e.Use(middleware.CORS())
	routes.Register(e)

	e.Logger.Fatal(e.Start(fmt.Sprintf(":%d", config.AppConf.Server.Http.Port)))
}

func (p *logicProgram) Stop() error {
	p.once.Do(func() {
		defer routes.CancelRoute(Echo)
		defer producer.LogicProducer.StopProducer()
		defer db.DisconnectMysql()
		defer db.DisconnectRedis()
	})
	return nil
}
