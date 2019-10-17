package rabbitMQ

import (
	"encoding/json"
	"github.com/heyHui2018/best-practise/base"
	"github.com/heyHui2018/best-practise/models/dataSource"
	"github.com/heyHui2018/log"
	"github.com/heyHui2018/utils"
	"github.com/streadway/amqp"
	"sync"
	"time"
)

var ConsumeWait *sync.WaitGroup
var MQCloseSign = false

func MQConsume() {
	if ConsumeChannel == nil {
		ConsumeConnect()
	}
	err := ConsumeChannel.Qos(1, 0, true)
	if err != nil {
		log.Warnf("MQConsume 设置Qos出错，err = %v", err)
		ConsumeChannel.Close()
		ConsumeChannel = nil
		ConsumeConn.Close()
		time.Sleep(3 * time.Second)
		go ConsumeStart()
		return
	}
	msg, err := ConsumeChannel.Consume(base.GetConfig().MQs["consume"].Queue, "", false, false, false, false, nil)
	if err != nil {
		log.Warn("MQConsume 接收MQ消息出错，err = ", err)
		ConsumeChannel.Close()
		ConsumeChannel = nil
		ConsumeConn.Close()
		time.Sleep(3 * time.Second)
		go ConsumeStart()
		return
	}

	ConsumeWait = new(sync.WaitGroup)
	for i := 0; i < base.GetConfig().MQs["consume"].ChanRangeNum; i++ {
		ConsumeWait.Add(1)
		go rangeChannel(msg)
	}
	ConsumeWait.Wait()

	if MQCloseSign == true {
		MQWait.Done()
	} else {
		log.Infof("MQCloseSign = %v,ConsumeQueue开始重连", MQCloseSign)
		ConsumeChannel.Close()
		ConsumeChannel = nil
		ConsumeConn.Close()
		go ConsumeStart()
	}
}

func rangeChannel(msg <-chan amqp.Delivery) {
	defer ConsumeWait.Done()
	for m := range msg {
		traceId := time.Now().Format("20060102150405") + utils.GetRandomString()
		data := new(dataSource.AirVisualReply)
		err := json.Unmarshal(m.Body, data)
		if err != nil {
			log.Warnf("rangeChannel Unmarshal msg 出错,traceId = %v,err = %v,msgInfo = %v", traceId, err, string(m.Body))
		} else {

		}
		m.Ack(false)
	}
}
