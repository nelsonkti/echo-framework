package main

import (
	"fmt"
	"github.com/judwhite/go-svc"
	"github.com/nelsonkti/echo-framework/config"
	"github.com/nelsonkti/echo-framework/lib/db/memcache"
	"github.com/nelsonkti/echo-framework/lib/db/mysql"
	"github.com/nelsonkti/echo-framework/lib/db/redis"
	"github.com/nelsonkti/echo-framework/lib/helper"
	"github.com/nelsonkti/echo-framework/lib/localtion"
	"github.com/nelsonkti/echo-framework/lib/logger"
	"github.com/nelsonkti/echo-framework/socket/nsq/consumer"
	"github.com/nelsonkti/echo-framework/socket/nsq/producer"
	"github.com/nelsonkti/echo-framework/socket/socketio_server"
	"github.com/nelsonkti/echo-framework/util/xetcd"
	"github.com/nelsonkti/echo-framework/util/xnsq"
	"github.com/nelsonkti/echo-framework/util/xnsq/service/registry"
	"log"
	"net/http"
	_ "net/http/pprof"
	"os"
	"path/filepath"
	"sync"
	"syscall"
)

type socketProgram struct {
	once sync.Once
}

func main() {
	p := &socketProgram{}
	if err := svc.Run(p, syscall.SIGINT, syscall.SIGTERM, syscall.SIGKILL); err != nil {
		logger.Sugar.Fatal(err)
	}
}

// svc 服务运行框架 程序启动时执行Init+Start, 服务终止时执行Stop
func (p *socketProgram) Init(env svc.Environment) error {
	if env.IsWindowsService() {
		dir := filepath.Dir(os.Args[0])
		return os.Chdir(dir)
	}
	return nil
}

func (p *socketProgram) Start() error {

	// 日志
	logger.New(
		logger.Base(config.AppConf.App.Env, config.AppConf.App.Path.LogPath),
		logger.FileName(helper.CurrentFileName()),
	)

	//性能分析
	go func() {
		log.Println(http.ListenAndServe("0.0.0.0:6060", nil))
	}()

	mysql.Connect()

	//连接 memcache
	memcache.Connect(config.AppConf.Data.Memcache.Host)

	//连接redis
	redis.NewClient(config.AppConf.Data.Redis.Addr, config.AppConf.Data.Redis.Password)

	xetcd.New(xetcd.Config{
		Endpoints: config.AppConf.Etcd.Host,
		OpenTLS:   config.AppConf.Etcd.OpenTls,
		TlsPath:   config.AppConf.Etcd.TlsPath,
		Env:       config.AppConf.App.Env,
		Pfx:       fmt.Sprintf("%s-%s", config.AppConf.App.Name, config.AppConf.App.Env),
	})

	// 启动消息队列服务
	go func() {
		defer helper.RecoverPanic()
		server := xnsq.NewNsqServer(registry.Options{
			NsqAddress:      config.AppConf.Mq.Nsq.Host,
			NSQConsumers:    config.AppConf.Mq.Nsq.Consumer,
			NSQAdminAddress: config.AppConf.Mq.Nsq.AdminAddress,
			Env:             config.AppConf.App.Env,
			LocalAddress:    localtion.GetLocalIP(),
		})
		server.Run(consumer.SocketConsumerHandler(server.Opt))
	}()

	go func() {
		socketio_server.Start(int(config.AppConf.Server.Socket.Port))
	}()

	return nil
}

func (p *socketProgram) Stop() error {
	p.once.Do(func() {
		defer mysql.Disconnect()
		defer redis.Disconnect()
		defer xetcd.Close()
		defer producer.SocketProducer.Stop()
		defer socketio_server.StopDevice()
	})
	return nil
}
