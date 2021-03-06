package socket_mq

import (
	"echo-framework/socket/socket_mq/producer"
)

var nsqAddress string
var address []string

//启动Nsq服务
func StartNsqServer(nsqAddr string, consumers []string, topicAddress string) {
	address = consumers
	nsqAddress = nsqAddr
	producer.StartNsqProducer(nsqAddr)
	startNsqConsumer(topicAddress)
}
