// Package producer
// @Author fuzengyao
// @Date 2022-11-09 11:15:43
package producer

import (
	"github.com/nelsonkti/echo-framework/util/xnsq/producer"
)

var Separator = "@"

var LogicProducer Producer

type Producer struct {
	producer.Producer
}

// 退出
func (l *Producer) Stop() {
	l.Producer.StopProducer()
}
