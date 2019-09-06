package kafka

import (
	"github.com/Shopify/sarama"
	"github.com/heyHui2018/best-practise/base"
	"github.com/heyHui2018/log"
	"strings"
)

func KafkaPublish() {
	config := sarama.NewConfig()
	config.Producer.Return.Successes = true                   // 是否等待成功后的响应
	config.Producer.RequiredAcks = sarama.WaitForAll          // 等待服务器所有副本都保存成功后的响应
	config.Producer.Partitioner = sarama.NewRandomPartitioner // 随机的分区类型

	msg := &sarama.ProducerMessage{}
	msg.Topic = base.GetConfig().Kafka.Topic
	msg.Partition = int32(-1)
	msg.Key = sarama.StringEncoder(base.GetConfig().Kafka.Key)
	msg.Value = sarama.ByteEncoder("hello,world")

	producer, err := sarama.NewSyncProducer(strings.Split(base.GetConfig().Kafka.Hosts, ","), config)
	if err != nil {
		log.Warnf("Failed to new producer,err = %v", err)
	}
	defer producer.Close()

	partition, offset, err := producer.SendMessage(msg)
	if err != nil {
		log.Warnf("Failed to produce message,err = %v", err)
	}
	log.Infof("partition = %v, offset = %v", partition, offset)
}
