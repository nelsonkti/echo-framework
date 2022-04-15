/**
** @创建时间 : 2021/11/13 16:03
** @作者 : fzy
 */
package xnsq

import (
	"github.com/nelsonkti/echo-framework/util/xnsq/producer"
	"github.com/nelsonkti/echo-framework/util/xnsq/server"
	"github.com/nelsonkti/echo-framework/util/xnsq/service/registry"
)

type NSQServer struct {
	Opt registry.Options
}

func NewNsqServer(opt registry.Options) NSQServer {
	return NSQServer{Opt: opt}
}

func (n *NSQServer) Run(c server.ConsumerHandler) (NSQServer *NSQServer) {
	n.startNsqProducer()
	n.startNsqConsumer(c)
	return
}

func (n *NSQServer) startNsqProducer() (NSQServer *NSQServer) {
	producer.StartNsqProducer(n.Opt)
	return
}

func (n *NSQServer) startNsqConsumer(c server.ConsumerHandler) (NSQServer *NSQServer) {
	c.Run()
	return
}
