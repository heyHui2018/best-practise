package kafka

import (
	"encoding/json"
	"github.com/Shopify/sarama"
	"github.com/heyHui2018/best-practise/base"
	"github.com/heyHui2018/log"
	"time"
)

var StopChan = make(chan struct{})
var StopPublishChan = make(chan struct{})
var MsgChan = make(chan string, 1000)

// 同步发送
func SyncProduce(log log.TLog, message string) {
	msg := &sarama.ProducerMessage{}
	msg.Topic = base.GetConfig().Kafka.PublishTopic
	msg.Key = sarama.StringEncoder(base.GetConfig().Kafka.Key)
	msg.Value = sarama.ByteEncoder(message)

	partition, offset, err := (*SyncProducer).SendMessage(msg)
	if err != nil {
		log.Warnf("SyncProduce,SendMessage error,err = %v", err)
		return
	}
	log.Infof("SyncProduce 完成,partition = %v, offset = %v", partition, offset)
}

// 异步发送
func AsyncProduce(log log.TLog) {
	go func(producer sarama.AsyncProducer) {
		for {
			select {
			case <-StopChan:
				StopPublishChan <- struct{}{}
				return
			case resp := <-producer.Successes():
				log.Infof("SyncProduce 完成,partition = %v, offset = %v", resp.Partition, resp.Offset)
			case resp := <-producer.Errors():
				log.Warnf("SyncProduce 失败,err = %v", resp.Err)
				time.Sleep(3 * time.Second)
				StopPublishChan <- struct{}{}
				go AsyncProduce(log)
				return
			}
		}
	}(*AsyncProducer)

	for {
		select {
		case <-StopPublishChan:
			return
		case msg := <-MsgChan:
			data, err := json.Marshal(msg)
			if err != nil {
				log.Warnf("AsyncProduce,Marshal error,err = %v", err)
				continue
			}
			sendMsg := &sarama.ProducerMessage{
				Topic: base.GetConfig().Kafka.PublishTopic,
				Key:   sarama.StringEncoder(base.GetConfig().Kafka.Key),
				Value: sarama.ByteEncoder(data),
			}
			(*AsyncProducer).Input() <- sendMsg
		}
	}

}
