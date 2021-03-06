package mq

import (
	"echo-framework/config"
	"fmt"
	"github.com/nsqio/go-nsq"
	"log"
	"os"
	"time"
)

// StartNsqConsume 启动nsq消费者，以后所有的消费者在这里注册
// 任何其他package不能import此package
func startNsqConsumer() {
	//通讯层发来的设备登录
	nsqConsumer("hello", "1", handleDeviceLogin, 100)
}

// nsqConsumer 消费消息
func nsqConsumer(topic, channel string, handle func(message *nsq.Message) error, concurrency int) {
	conf := nsq.NewConfig()
	conf.LookupdPollInterval = 1 * time.Second
	conf.MaxInFlight = 10 + len(config.NSQServerHosts)

	consumer, err := nsq.NewConsumer(topic, channel, conf)
	if err != nil {
		panic(err)
	}
	consumer.AddConcurrentHandlers(nsq.HandlerFunc(handle), concurrency)
	consumer.SetLogger(log.New(os.Stderr, "", log.Flags()), nsq.LogLevelError)

	if config.Env == "local" {
		err = consumer.ConnectToNSQD(nsqAddress)
	} else {
		err = consumer.ConnectToNSQLookupds(address)
	}
	if err != nil {
		panic(err)
	}
}

//  设备上线
func handleDeviceLogin(msg *nsq.Message) error {
	msg.Finish()
	data := string(msg.Body)
	//logger.Sugar.Info(data)

	fmt.Println(data)

	return nil
}
