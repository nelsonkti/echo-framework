package mq

import "echo-framework/logic/mq/producer"

var address []string
var nsqAddress string

//启动Nsq服务
func StartNsqServer(addr string, consumers []string) {
	address = consumers
	nsqAddress = addr
	producer.StartNsqProducer(addr)
	startNsqConsumer()
}
