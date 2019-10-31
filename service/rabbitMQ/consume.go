package rabbitMQ

import (
	"encoding/json"
	"sync"
	"time"

	"github.com/heyHui2018/log"
	"github.com/heyHui2018/utils"
	"github.com/streadway/amqp"

	"github.com/heyHui2018/best-practise/base"
	"github.com/heyHui2018/best-practise/model/dataSource"
)

var wait sync.WaitGroup

func Consume() {
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

	for i := 0; i < base.GetConfig().MQs["consume"].ChanRangeNum; i++ {
		wait.Add(1)
		go rangeChannel(msg)
	}
	wait.Wait()

	if MQCloseSign {
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
	defer wait.Done()
	for m := range msg {
		traceId := time.Now().Format("20060102150405") + utils.GetRandomString()
		data := new(dataSource.AirVisualReply)
		err := json.Unmarshal(m.Body, data)
		if err != nil {
			log.Warnf("rangeChannel Unmarshal msg 出错,traceId = %v,err = %v,msgInfo = %v", traceId, err, string(m.Body))
			continue
		}
		// do something with data
		m.Ack(false)
	}
}
