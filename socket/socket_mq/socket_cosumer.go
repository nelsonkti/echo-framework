package socket_mq

import (
	"echo-framework/config"
	"echo-framework/lib/logger"
	"github.com/nsqio/go-nsq"
	"log"
	"os"
	"time"
)

func startNsqConsumer(localAddr string) {

}

// nsqConsumer 消费消息
func nsqConsumer(topic, channel string, handle func(message *nsq.Message) error, concurrency int) {

	conf := nsq.NewConfig()
	conf.LookupdPollInterval = 1 * time.Second
	conf.MaxInFlight = 1 + len(config.NSQServerHosts)

	consumer, err := nsq.NewConsumer(topic, channel, conf)
	if err != nil {
		logger.Sugar.Error(err)
		panic(err)
	}
	consumer.AddConcurrentHandlers(nsq.HandlerFunc(handle), concurrency)
	consumer.SetLogger(log.New(os.Stderr, "", log.Flags()), nsq.LogLevelError)

	err = consumer.ConnectToNSQD(nsqAddress)
	if err != nil {
		panic(err)
	}
	if config.Env != "local" {
		err = consumer.ConnectToNSQLookupds(address)
		if err != nil {
			panic(err)
		}
	}
}
