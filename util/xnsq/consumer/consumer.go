// Package consumer
// @Author fuzengyao
// @Date 2022-11-10 11:07:56
package consumer

import (
	"github.com/nelsonkti/echo-framework/util/xnsq/service/registry"
	"github.com/nsqio/go-nsq"
	"log"
	"os"
	"time"
)

var Options registry.Options

// NsqConsumer
// @Description: 消费消息
// @param topic
// @param channel
// @param handle
// @param concurrency
func NsqConsumer(topic, channel string, handle func(message *nsq.Message) error, concurrency int) {
	conf := nsq.NewConfig()
	conf.LookupdPollInterval = 1 * time.Second
	conf.MaxInFlight = 1000

	consumer, err := nsq.NewConsumer(topic, channel, conf)
	if err != nil {
		panic(err)
	}
	consumer.AddConcurrentHandlers(nsq.HandlerFunc(handle), concurrency)
	consumer.SetLogger(log.New(os.Stderr, "", log.Flags()), nsq.LogLevelError)

	err = consumer.ConnectToNSQD(Options.NsqAddress)
	if err != nil {
		panic(err)
	}
	if Options.Env != "local" {
		err = consumer.ConnectToNSQLookupds(Options.NSQConsumers)
		if err != nil {
			panic(err)
		}
	}
}
