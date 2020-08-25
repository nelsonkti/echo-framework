package producer

import (
    "github.com/nsqio/go-nsq"
    "echo-framework/lib/logger"
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
        panic("nsq new panic")
    }

    err = producer.Ping()
    if nil != err {
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

func StopProducer() {
    if producer != nil {
        producer.Stop()
    }
    logger.Sugar.Info("stop nsq producer")
}
