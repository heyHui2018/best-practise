package rabbitMQ

import (
	"encoding/json"
	"time"

	"github.com/heyHui2018/log"
	"github.com/streadway/amqp"

	"github.com/heyHui2018/best-practise/base"
)

func MQPublish(t log.TLog, msg interface{}) {
	if PublishChannel == nil {
		PublishStart()
	}
	publishInfo, err := json.Marshal(msg)
	if err != nil {
		t.Warnf("MQPublish Marshal error,err = %v", err)
		return
	}
	exchange := base.GetConfig().MQs["publish"].Exchange
	// exchange := base.GetConfig().MQs["deadLetter"].Exchange
	key := base.GetConfig().MQs["publish"].Key
	err = PublishChannel.Publish(exchange, key, false, false, amqp.Publishing{
		ContentType: "text/plain",
		// DeliveryMode: 2, // exchange/queue均为durable时，在此设置DeliveryMode为2即可持久化数据
		Body: publishInfo,
	})
	if err != nil {
		t.Warnf("MQPublish error,err = %v", err)
		PublishChannel.Close()
		PublishChannel = nil
		PublishConn.Close()
		time.Sleep(3 * time.Second)
		go PublishStart()
		return
	}
	t.Infof("MQPublish success")
}
