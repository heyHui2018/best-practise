package service

import (
	"encoding/json"
	"github.com/heyHui2018/best-practise/base"
	"github.com/ngaut/log"
	"github.com/streadway/amqp"
	"time"
)

func MQPublish(traceId string, msg interface{}) {
	if PublishChannel == nil {
		PublishStart()
	}
	publishInfo, err := json.Marshal(msg)
	if err != nil {
		log.Warnf("MQPublish Marshal error,traceId = %v,err = %v", traceId, err)
		return
	}
	exchange := base.GetConfig().MQs["publish"].Exchange
	key := base.GetConfig().MQs["publish"].Key
	err = PublishChannel.Publish(exchange, key, false, false, amqp.Publishing{
		ContentType: "text/plain",
		// DeliveryMode: 2, // exchange/queue均为durable时，在此设置DeliveryMode为2即可持久化数据
		Body: publishInfo,
	})
	if err != nil {
		log.Warnf("MQPublish error,traceId = %v,err = %v", traceId, err)
		PublishChannel.Close()
		PublishChannel = nil
		PublishConn.Close()
		time.Sleep(3 * time.Second)
		go PublishStart()
		return
	}
	log.Infof("MQPublish success,traceId = %v", traceId)
}
