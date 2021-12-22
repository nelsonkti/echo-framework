/**
** @创建时间 : 2021/11/15 15:14
** @作者 : fzy
 */
package consumer

import (
	"echo-framework/util/xnsq/consumer"
	"echo-framework/util/xnsq/server"
	"echo-framework/util/xnsq/service/registry"
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
