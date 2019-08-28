package service

import (
	"github.com/heyHui2018/best-practise/base"
	"github.com/ngaut/log"
	"sync"
	"time"
)

func MQStart() {
	go ConsumeStart()
	go PublishStart()
}

func ConsumeStart() {
	base.ConsumeConnect()
	exchange := base.GetConfig().MQs["consume"].Exchange
	err := base.ConsumeChannel.ExchangeDeclare(exchange, "direct", true, false, false, false, nil)
	if nil != err {
		log.Warnf("consumeStart 初始化 Exchange:%v 出错,err = %v", exchange, err)
		base.ConsumeChannel.Close()
		base.ConsumeConn.Close()
		time.Sleep(3 * time.Second)
		go ConsumeStart()
		return
	}
	queue := base.GetConfig().MQs["consume"].Queue
	_, err = base.ConsumeChannel.QueueDeclare(queue, true, false, false, false, nil)
	if err != nil {
		log.Warnf("consumeStart 初始化 Queue:%v 出错,err = %v", queue, err)
		base.ConsumeChannel.Close()
		base.ConsumeConn.Close()
		time.Sleep(3 * time.Second)
		go ConsumeStart()
		return
	}
	key := base.GetConfig().MQs["consume"].Key
	err = base.ConsumeChannel.QueueBind(queue, key, exchange, false, nil)
	if err != nil {
		log.Warnf("consumeStart 绑定 %v 到 %v 出错,err = %v", queue, exchange, err)
		base.ConsumeChannel.Close()
		base.ConsumeConn.Close()
		time.Sleep(3 * time.Second)
		go ConsumeStart()
		return
	}
	log.Info("consumeStart完成")
	base.MQWait = new(sync.WaitGroup)
	base.MQWait.Add(1)
	go Consume()
}

func PublishStart() {
	base.PublishConnect()
	exchange := base.GetConfig().MQs["publish"].Exchange
	err := base.PublishChannel.ExchangeDeclare(exchange, "topic", true, false, false, false, nil)
	if nil != err {
		log.Warnf("publishStart 初始化 Exchange:%v 出错,err = %v", exchange, err)
		base.PublishChannel.Close()
		base.PublishConn.Close()
		time.Sleep(3 * time.Second)
		go PublishStart()
		return
	}
	queue := base.GetConfig().MQs["publish"].Queue
	_, err = base.PublishChannel.QueueDeclare(queue, true, false, false, false, nil)
	if err != nil {
		log.Warnf("publishStart 初始化 Queue:%v 出错,err = %v", queue, err)
		base.PublishChannel.Close()
		base.PublishConn.Close()
		time.Sleep(3 * time.Second)
		go PublishStart()
		return
	}
	key := base.GetConfig().MQs["publish"].Key
	err = base.PublishChannel.QueueBind(queue, key, exchange, false, nil)
	if err != nil {
		log.Warnf("publishStart 绑定 %v 到 %v 出错,err = %v", queue, exchange, err)
		base.PublishChannel.Close()
		base.PublishConn.Close()
		time.Sleep(3 * time.Second)
		go PublishStart()
		return
	}
	log.Info("publishStart完成")
}
