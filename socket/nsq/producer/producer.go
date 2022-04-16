/**
** @创建时间 : 2021/11/15 15:10
** @作者 : fzy
 */
package producer

import "github.com/nelsonkti/echo-framework/util/xnsq/producer"

var Separator = "@"

var SocketProducer Producer

type Producer struct {
	producer.Producer
}

// 退出
func (l *Producer) Stop() {
	l.Producer.StopProducer()
}
