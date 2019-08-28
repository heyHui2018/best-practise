package service

import (
	"encoding/json"
	"github.com/heyHui2018/best-practise/base"
	"github.com/ngaut/log"
	"github.com/streadway/amqp"
	"time"
)

func Publish(traceId string, msg interface{}) {
	if base.PublishChannel == nil {
		PublishStart()
	}
	publishInfo, err := json.Marshal(msg)
	if err != nil {
		log.Warnf("Publish Marshal error,traceId = %v,err = %v", traceId, err)
		return
	}
	exchange := base.GetConfig().MQs["publish"].Exchange
	key := base.GetConfig().MQs["publish"].Key
	err = base.PublishChannel.Publish(exchange, key, false, false, amqp.Publishing{
		ContentType: "text/plain",
		// DeliveryMode: 2, // exchange/queue均为durable时，在此设置DeliveryMode为2即可持久化数据
		Body: publishInfo,
	})
	if err != nil {
		log.Warnf("Publish error,traceId = %v,err = %v", traceId, err)
		base.PublishChannel.Close()
		base.PublishChannel = nil
		base.PublishConn.Close()
		time.Sleep(3 * time.Second)
		go PublishStart()
		return
	}
	log.Infof("Publish success,traceId = %v", traceId)
}
