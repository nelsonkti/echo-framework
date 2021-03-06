package main

import (
	"echo-framework/config"
	"echo-framework/lib/db"
	"echo-framework/lib/helper"
	"echo-framework/lib/logger"
	"echo-framework/socket/socket_mq"
	"echo-framework/socket/socketio_server"
	_ "github.com/judwhite/go-svc"
	"github.com/judwhite/go-svc/svc"
	"log"
	"net/http"
	_ "net/http/pprof"
	"os"
	"path/filepath"
	"strconv"
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
	//性能分析
	go func() {
		log.Println(http.ListenAndServe("0.0.0.0:6060", nil))
	}()

	db.InitMysql()
	// 连接redis
	db.ConnectRedis(config.RedisIP, config.RedisPassword, 0, "default")

	// 启动消息队列服务
	go func() {
		defer helper.RecoverPanic()

		socket_mq.StartNsqServer(config.NSQIP, config.NSQConsumers, helper.GetLocalIP())
	}()

	go func() {
		port, _ := strconv.Atoi(config.ConnectTCPListenPort)
		socketio_server.Start(port)
	}()

	return nil
}

func (p *socketProgram) Stop() error {
	p.once.Do(func() {
		defer socketio_server.StopDevice()
	})
	return nil
}
