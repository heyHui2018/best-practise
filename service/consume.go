package service

import (
	"encoding/json"
	"github.com/heyHui2018/best-practise/base"
	"github.com/heyHui2018/best-practise/models"
	"github.com/heyHui2018/utils"
	"github.com/ngaut/log"
	"github.com/streadway/amqp"
	"sync"
	"time"
)

var ConsumeWait *sync.WaitGroup
var ConsumeRegularCloseSign = false

func Consume() {
	if base.ConsumeChannel == nil {
		base.ConsumeConnect()
	}
	err := base.ConsumeChannel.Qos(1, 0, true)
	if err != nil {
		log.Warnf("ConsumeBindQueue 设置Qos出错，err = %v", err)
		base.ConsumeChannel.Close()
		base.ConsumeChannel = nil
		base.ConsumeConn.Close()
		time.Sleep(3 * time.Second)
		go ConsumeStart()
		return
	}
	msg, err := base.ConsumeChannel.Consume(base.GetConfig().MQs["consume"].Queue, "", false, false, false, false, nil)
	if err != nil {
		log.Warn("ConsumeBindQueue 接收MQ消息出错，err = ", err)
		base.ConsumeChannel.Close()
		base.ConsumeChannel = nil
		base.ConsumeConn.Close()
		time.Sleep(3 * time.Second)
		go ConsumeStart()
		return
	}

	ConsumeWait = new(sync.WaitGroup)
	for i := 0; i < base.GetConfig().MQs["consume"].ChanRangeNum; i++ {
		ConsumeWait.Add(1)
		go rangeBindChannel(msg)
	}
	ConsumeWait.Wait()

	if ConsumeRegularCloseSign == true {
		log.Infof("ConsumeRegularCloseSign = %v,ConsumeQueue已正常关闭", ConsumeRegularCloseSign)
		base.MQWait.Done()
	} else {
		log.Infof("ConsumeRegularCloseSign = %v,ConsumeQueue开始重连", ConsumeRegularCloseSign)
		base.ConsumeChannel.Close()
		base.ConsumeChannel = nil
		base.ConsumeConn.Close()
		go ConsumeStart()
	}
}

func rangeBindChannel(msg <-chan amqp.Delivery) {
	defer ConsumeWait.Done()
	for m := range msg {
		traceId := time.Now().Format("20060102150405") + utils.GetRandomString()
		data := new(models.BindingMQData)
		err := json.Unmarshal(m.Body, data)
		if err != nil {
			log.Warnf("rangeBindChannel Unmarshal msg 出错,traceId = %v,err = %v,msgInfo = %v", traceId, err, string(m.Body))
		} else {

		}
		m.Ack(false)
	}
}
