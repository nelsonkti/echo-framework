/**
** @创建时间 : 2021/11/15 15:16
** @作者 : fzy
 */
package consumer

import (
	"echo-framework/util/xnsq/consumer"
	"echo-framework/util/xnsq/server"
	"echo-framework/util/xnsq/service/registry"
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
