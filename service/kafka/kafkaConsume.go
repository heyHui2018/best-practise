package kafka

import (
	"github.com/Shopify/sarama"
	cluster "github.com/bsm/sarama-cluster"
	"github.com/heyHui2018/best-practise/base"
	"github.com/heyHui2018/log"
	"strings"
	"time"
)

// "github.com/Shopify/sarama"不支持consume group,故选择再次封装的"github.com/bsm/sarama-cluster"
func Consume(log log.TLog, groupId string) {
	config := cluster.NewConfig()
	config.Consumer.Return.Errors = true
	config.Consumer.Offsets.CommitInterval = 1 * time.Second
	config.Consumer.Offsets.Initial = sarama.OffsetNewest // 从最新的offset开始
	config.Group.Return.Notifications = true

	// 第二个参数是groupId
	consumer, err := cluster.NewConsumer(strings.Split(base.GetConfig().Kafka.Hosts, ","), groupId, strings.Split(base.GetConfig().Kafka.Topic, ","), config)
	if err != nil {
		log.Warnf("KafkaConsume NewConsumer error,err = %v", err)
	}
	defer consumer.Close()

	// 接收错误
	go func() {
		for err := range consumer.Errors() {
			log.Warnf("KafkaConsume,error = %v", err)
		}
	}()

	// 接收rebalance信息
	go func() {
		for ntf := range consumer.Notifications() {
			log.Warnf("KafkaConsume，Notification = %+v", ntf)
		}
	}()

	// 消费消息
	for msg := range consumer.Messages() {
		log.Infof("KafkaConsume,topic = %v,partition = %v,offset = %v,key = %v,value = %v", msg.Topic, msg.Partition, msg.Offset, msg.Key, msg.Value)
		consumer.MarkOffset(msg, "") // 提交offset
		// if err := consumer.CommitOffsets(); err != nil {
		// 	panic(err)
		// }
	}
}
