package producer

import (
	"echo-framework/lib/logger"
	"github.com/nsqio/go-nsq"
	"time"
)

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

func DeferredPublish(deviceId, serverAddress, topic, data string, delay time.Duration) {
	if serverAddress == "" {
		//logger.Sugar.Infow("device not online:", "device_id:", deviceId, "topic:", topic)
		return
	}

	err := producer.DeferredPublish(serverAddress+"."+topic, delay, []byte(deviceId+Separator+data))
	if err != nil {
		logger.Sugar.Error(err)
	}
}

// 指定服务发布
func AssignServerPublish(serverAddress, topic, data string) {
	err := producer.Publish(serverAddress+"."+topic, []byte(Separator+data))
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

func StopProducer() {
	if producer != nil {
		producer.Stop()
	}
	logger.Sugar.Info("stop nsq producer")
}
