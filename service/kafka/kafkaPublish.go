package kafka

import (
	"encoding/json"
	"github.com/Shopify/sarama"
	"github.com/heyHui2018/best-practise/base"
	"github.com/heyHui2018/log"
	"strings"
	"time"
)

var StopChan = make(chan struct{})
var StopPublishChan = make(chan struct{})
var MsgChan = make(chan string, 1000)

func config() *sarama.Config {
	config := sarama.NewConfig()
	config.Producer.Return.Successes = true // 是否等待成功后的响应
	config.Producer.Return.Errors = true
	config.Producer.Timeout = 10 * time.Second
	config.Producer.RequiredAcks = sarama.WaitForAll          // 等待服务器所有副本都保存成功后的响应
	config.Producer.Partitioner = sarama.NewRandomPartitioner // 随机的分区类型
	config.Version = sarama.V0_10_0_1                         // 版本设置有问题时,kafka会无法发送消息
	return config
}

// 同步发送
func SyncProduce(log log.TLog, message string) {
	msg := &sarama.ProducerMessage{}
	msg.Topic = base.GetConfig().Kafka.Topic
	msg.Key = sarama.StringEncoder(base.GetConfig().Kafka.Key)
	msg.Value = sarama.ByteEncoder(message)

	producer, err := sarama.NewSyncProducer(strings.Split(base.GetConfig().Kafka.Hosts, ","), config())
	if err != nil {
		log.Warnf("SyncProduce,NewSyncProducer error,err = %v", err)
		return
	}
	defer producer.Close()

	partition, offset, err := producer.SendMessage(msg)
	if err != nil {
		log.Warnf("SyncProduce,SendMessage error,err = %v", err)
		return
	}
	log.Infof("SyncProduce 完成,partition = %v, offset = %v", partition, offset)
}

// 异步发送
func AsyncProduce(log log.TLog) {
	producer, err := sarama.NewAsyncProducer(strings.Split(base.GetConfig().Kafka.Hosts, ","), config())
	if err != nil {
		log.Warnf("AsyncProduce,NewSyncProducer error,err = %v", err)
		return
	}
	defer producer.Close()

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
	}(producer)

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
				Topic: base.GetConfig().Kafka.Topic,
				Key:   sarama.StringEncoder(base.GetConfig().Kafka.Key),
				Value: sarama.ByteEncoder(data),
			}
			producer.Input() <- sendMsg
		}
	}

}
