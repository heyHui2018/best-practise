package rabbitMQ

import (
	"github.com/heyHui2018/best-practise/base"
	"github.com/heyHui2018/log"
	"github.com/streadway/amqp"
	"sync"
	"time"
)

var MQWait sync.WaitGroup
var MQCloseSign = false
var ConsumeConn *amqp.Connection
var ConsumeChannel *amqp.Channel
var PublishConn *amqp.Connection
var PublishChannel *amqp.Channel

func MQInit() {
	MQWait.Add(1)
	go ConsumeStart()
	go PublishStart()
}

func ConsumeStart() {
	if MQCloseSign {
		return
	}
	ConsumeConnect()
	exchange := base.GetConfig().MQs["consume"].Exchange
	err := ConsumeChannel.ExchangeDeclare(exchange, "direct", true, false, false, false, nil)
	if nil != err {
		log.Warnf("consumeStart 初始化 Exchange:%v 出错,err = %v", exchange, err)
		ConsumeChannel.Close()
		ConsumeConn.Close()
		time.Sleep(3 * time.Second)
		go ConsumeStart()
		return
	}
	queue := base.GetConfig().MQs["consume"].Queue
	_, err = ConsumeChannel.QueueDeclare(queue, true, false, false, false, nil)
	if err != nil {
		log.Warnf("consumeStart 初始化 Queue:%v 出错,err = %v", queue, err)
		ConsumeChannel.Close()
		ConsumeConn.Close()
		time.Sleep(3 * time.Second)
		go ConsumeStart()
		return
	}
	key := base.GetConfig().MQs["consume"].Key
	err = ConsumeChannel.QueueBind(queue, key, exchange, false, nil)
	if err != nil {
		log.Warnf("consumeStart 绑定 %v 到 %v 出错,err = %v", queue, exchange, err)
		ConsumeChannel.Close()
		ConsumeConn.Close()
		time.Sleep(3 * time.Second)
		go ConsumeStart()
		return
	}
	log.Info("consumeStart完成")
	go Consume()
}

func PublishStart() {
	PublishConnect()
	exchange := base.GetConfig().MQs["publish"].Exchange
	err := PublishChannel.ExchangeDeclare(exchange, "topic", true, false, false, false, nil)
	if nil != err {
		log.Warnf("publishStart 初始化 Exchange:%v 出错,err = %v", exchange, err)
		PublishChannel.Close()
		PublishConn.Close()
		time.Sleep(3 * time.Second)
		go PublishStart()
		return
	}
	queue := base.GetConfig().MQs["publish"].Queue
	_, err = PublishChannel.QueueDeclare(queue, true, false, false, false, nil)
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
	log.Info("publishStart完成")
}

func ConsumeConnect() {
	log.Info("ConsumeConnect 开始连接")
	var err error
	username := base.GetConfig().MQs["consume"].Username
	password := base.GetConfig().MQs["consume"].Password
	ip := base.GetConfig().MQs["consume"].Ip
	port := base.GetConfig().MQs["consume"].Port
	host := base.GetConfig().MQs["consume"].Host
	mqUrl := "amqp://" + username + ":" + password + "@" + ip + ":" + port + "/" + host
a:
	ConsumeConn, err = amqp.Dial(mqUrl)
	if err != nil {
		log.Warnf("ConsumeConnect 连接MQ出错，err = %v", err)
		time.Sleep(3 * time.Second)
		goto a
	}
	ConsumeChannel, err = ConsumeConn.Channel()
	if err != nil {
		log.Warnf("ConsumeConnect 打开channel出错，err = %v", err)
		ConsumeConn.Close()
		time.Sleep(3 * time.Second)
		goto a
	}
	log.Info("ConsumeConnect 连接完成")
}

func PublishConnect() {
	log.Info("PublishConnect 开始连接")
	var err error
	username := base.GetConfig().MQs["publish"].Username
	password := base.GetConfig().MQs["publish"].Password
	ip := base.GetConfig().MQs["publish"].Ip
	port := base.GetConfig().MQs["publish"].Port
	host := base.GetConfig().MQs["publish"].Host
	mqUrl := "amqp://" + username + ":" + password + "@" + ip + ":" + port + "/" + host
a:
	PublishConn, err = amqp.Dial(mqUrl)
	if err != nil {
		log.Warnf("PublishConnect 连接MQ出错，err = %v", err)
		time.Sleep(3 * time.Second)
		goto a
	}
	PublishChannel, err = PublishConn.Channel()
	if err != nil {
		log.Warnf("PublishConnect 打开channel出错，err = %v", err)
		PublishConn.Close()
		time.Sleep(3 * time.Second)
		goto a
	}
	log.Info("PublishConnect 连接完成")
}
