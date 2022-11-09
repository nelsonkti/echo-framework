// Package consumer
// @Author fuzengyao
// @Date 2022-11-09 11:16:20
package consumer

import (
	"github.com/nelsonkti/echo-framework/util/xnsq/consumer"
	"github.com/nelsonkti/echo-framework/util/xnsq/server"
	"github.com/nelsonkti/echo-framework/util/xnsq/service/registry"
)

func SocketConsumerHandler(opt registry.Options) server.ConsumerHandler {
	consumer.Options = opt
	return &SocketNsqConsumer{Options: opt}
}

type SocketNsqConsumer struct {
	registry.Options
}

func (l *SocketNsqConsumer) Run() {

}
