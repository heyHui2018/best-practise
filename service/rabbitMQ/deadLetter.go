package rabbitMQ

import (
	"time"

	"github.com/heyHui2018/log"
	"github.com/streadway/amqp"

	"github.com/heyHui2018/best-practise/base"
)

func DeadLetterStart() {
	PublishConnect()
	// 死信
	deadExchange := base.GetConfig().MQs["deadLetter"].Exchange
	err := PublishChannel.ExchangeDeclare(deadExchange, "direct", true, false, false, false, nil)
	if err != nil {
		log.Warnf("DeadLetterStart,deadExchange 初始化 出错,err = %v", err)
		PublishChannel.Close()
		PublishConn.Close()
		time.Sleep(3 * time.Second)
		go DeadLetterStart()
		return
	}
	deadQueue := base.GetConfig().MQs["deadLetter"].Queue
	_, err = PublishChannel.QueueDeclare(deadQueue, true, false, false, false, nil)
	if err != nil {
		log.Warnf("DeadLetterStart,deadQueue 初始化 出错,err = %v", err)
		PublishChannel.Close()
		PublishConn.Close()
		time.Sleep(3 * time.Second)
		go DeadLetterStart()
		return
	}
	deadRoutingKey := base.GetConfig().MQs["deadLetter"].Key
	err = PublishChannel.QueueBind(deadQueue, deadRoutingKey, deadExchange, false, nil)
	if err != nil {
		log.Warnf("DeadLetterStart,绑定 %v 到 %v 出错,err = %v", deadQueue, deadExchange, err)
		PublishChannel.Close()
		PublishConn.Close()
		time.Sleep(3 * time.Second)
		go DeadLetterStart()
		return
	}
	// normal
	exchange := base.GetConfig().MQs["publish"].Exchange
	err = PublishChannel.ExchangeDeclare(exchange, "topic", true, false, false, false, nil)
	if nil != err {
		log.Warnf("publishStart 初始化 Exchange:%v 出错,err = %v", exchange, err)
		PublishChannel.Close()
		PublishConn.Close()
		time.Sleep(3 * time.Second)
		go PublishStart()
		return
	}
	args := amqp.Table{}
	args["x-dead-letter-exchange"] = deadExchange
	args["x-dead-letter-routing-key"] = deadRoutingKey
	queue := base.GetConfig().MQs["publish"].Queue
	_, err = PublishChannel.QueueDeclare(queue, true, false, false, false, args)
	if err != nil {
		log.Warnf("publishStart 初始化 Queue:%v 出错,err = %v", queue, err)
		PublishChannel.Close()
		PublishConn.Close()
		time.Sleep(3 * time.Second)
		go PublishStart()
		return
	}
	key := base.GetConfig().MQs["publish"].Key
	err = PublishChannel.QueueBind(queue, key, exchange, false, nil)
	if err != nil {
		log.Warnf("publishStart 绑定 %v 到 %v 出错,err = %v", queue, exchange, err)
		PublishChannel.Close()
		PublishConn.Close()
		time.Sleep(3 * time.Second)
		go PublishStart()
		return
	}
}
