package kafka

import (
	"sync"
	"time"

	cluster "github.com/bsm/sarama-cluster"
	"github.com/heyHui2018/log"
	"github.com/heyHui2018/utils"
)

var wait sync.WaitGroup

func Consume(consumer *cluster.Consumer) {
	defer wait.Done()
	for msg := range consumer.Messages() {
		t := new(log.TLog)
		t.TraceId = time.Now().Format("20060102150405") + utils.GetRandomString()
		t.Infof("KafkaConsume,topic = %v,partition = %v,offset = %v,key = %v,value = %v", msg.Topic, msg.Partition, msg.Offset, msg.Key, msg.Value)
		consumer.MarkOffset(msg, "") // 提交offset
		// if err := consumer.CommitOffsets(); err != nil {
		// 	panic(err)
		// }
	}
}
