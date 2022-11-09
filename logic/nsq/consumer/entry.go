// Package consumer
// @Author fuzengyao
// @Date 2022-11-09 11:15:37
package consumer

import (
	"github.com/nelsonkti/echo-framework/util/xnsq/consumer"
	"github.com/nelsonkti/echo-framework/util/xnsq/server"
	"github.com/nelsonkti/echo-framework/util/xnsq/service/registry"
)

func LogicConsumerHandler(opt registry.Options) server.ConsumerHandler {
	consumer.Options = opt
	return &LogicNsqConsumer{opt}
}

type LogicNsqConsumer struct {
	registry.Options
}

func (l *LogicNsqConsumer) Run() {

}
