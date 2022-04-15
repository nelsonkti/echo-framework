/**
** @创建时间 : 2021/11/15 15:14
** @作者 : fzy
 */
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
