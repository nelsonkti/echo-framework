package producer

import (
	"echo-framework/lib/logger"
	"github.com/nsqio/go-nsq"
)

//nsq 服务 生产者
var producer *nsq.Producer
var Separator = "@"

func StartNsqProducer(addr string) {
	if producer != nil {
		return
	}
	var err error
	cfg := nsq.NewConfig()
	producer, err = nsq.NewProducer(addr, cfg)
	if nil != err {
		logger.Sugar.Info(err)
		panic("nsq new panic")
	}

	err = producer.Ping()
	if nil != err {
		logger.Sugar.Info(err)
		panic("nsq ping panic")
	}
}

func Publish(topic, data string) {
	err := producer.Publish(topic, []byte(Separator+data))
	if err != nil {
		logger.Sugar.Error(err)
	}
}

func AssignUuidPublish(uuid, topic, data string) {
	if uuid == "" {
		logger.Sugar.Error("uuid为空")
		return
	}
	err := producer.Publish(topic, []byte(uuid+Separator+data))
	if err != nil {
		logger.Sugar.Error(err)
	}
}
