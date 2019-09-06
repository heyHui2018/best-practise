package kafka

import (
	"github.com/Shopify/sarama"
	"github.com/heyHui2018/best-practise/base"
	"github.com/heyHui2018/log"
	"strings"
	"sync"
)

var wg sync.WaitGroup

func kafkaConsume() {
	// 连接kafka
	consumer, err := sarama.NewConsumer(strings.Split(base.GetConfig().Kafka.Hosts, ","), nil)
	if err != nil {
		log.Warnf("Failed to new consumer,err = %v", err)
		return
	}
	defer consumer.Close()

	// consumer.Partitions 用户获取Topic上所有的Partitions
	partitionList, err := consumer.Partitions(base.GetConfig().Kafka.Topic)
	if err != nil {
		log.Warnf("Failed to get the list of partitions,err = %v", err)
		return
	}

	for partition := range partitionList {
		pc, err := consumer.ConsumePartition(base.GetConfig().Kafka.Topic, int32(partition), sarama.OffsetNewest)
		if err != nil {
			log.Warnf("Failed to consume for partition = %v,err = %v", partition, err)
		}
		pc.AsyncClose()
		wg.Add(1)
		go func(sarama.PartitionConsumer) {
			defer wg.Done()
			for msg := range pc.Messages() {
				log.Infof("Partition:%d, Offset:%d, Key:%s, Value:%s", msg.Partition, msg.Offset, string(msg.Key), string(msg.Value))
			}
		}(pc)
	}
	wg.Wait()
	log.Infof("Done consuming topic 'test'")
}
