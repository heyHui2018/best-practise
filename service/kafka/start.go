package kafka

import (
	"github.com/Shopify/sarama"
	cluster "github.com/bsm/sarama-cluster"
	"github.com/heyHui2018/best-practise/base"
	"github.com/heyHui2018/log"
	"sync"
	"time"
)

var KafkaWait sync.WaitGroup
var KafkaCloseSign = false
var Consumer *cluster.Consumer
var SyncProducer *sarama.SyncProducer
var AsyncProducer *sarama.AsyncProducer

func KafkaInit() {
	SyncProducer = new(sarama.SyncProducer)
	AsyncProducer = new(sarama.AsyncProducer)
	go SyncPublishStart()
	go AsyncPublishStart()
	KafkaWait.Add(1)
	go ConsumeStart()
}

// "github.com/Shopify/sarama"不支持consume group,故选择再次封装的"github.com/bsm/sarama-cluster"
func ConsumeStart() {
	if KafkaCloseSign {
		return
	}
	config := cluster.NewConfig()
	config.Consumer.Return.Errors = true
	config.Consumer.Offsets.CommitInterval = 1 * time.Second
	config.Consumer.Offsets.Initial = sarama.OffsetNewest // 从最新的offset开始
	config.Group.Return.Notifications = true

	var err error
	// groupId:自定义字符串,多个consumer设置相同的groupId时,相当于多消费者消费一个队列
	Consumer, err = cluster.NewConsumer(base.GetConfig().Kafka.Hosts, base.GetConfig().Kafka.ConsumeGroupId, base.GetConfig().Kafka.ConsumeTopic, config)
	if err != nil {
		log.Warnf("KafkaConsume NewConsumer error,err = %v", err)
		time.Sleep(3 * time.Second)
		go ConsumeStart()
		return
	}

	// 接收错误
	go func() {
		for err := range Consumer.Errors() {
			log.Warnf("KafkaConsume,error = %v", err)
		}
	}()

	// 接收rebalance信息
	go func() {
		for ntf := range Consumer.Notifications() {
			log.Warnf("KafkaConsume，Notification = %+v", ntf)
		}
	}()

	for i := 0; i < base.GetConfig().MQs["consume"].ChanRangeNum; i++ {
		wait.Add(1)
		go Consume(Consumer)
	}
	wait.Wait()

	if KafkaCloseSign {
		KafkaWait.Done()
	} else {
		log.Infof("KafkaCloseSign = %v,KafkaConsume开始重连", KafkaCloseSign)
		go ConsumeStart()
	}
}

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
func SyncPublishStart() {
	if KafkaCloseSign {
		return
	}
	var err error
	*SyncProducer, err = sarama.NewSyncProducer(base.GetConfig().Kafka.Hosts, config())
	if err != nil {
		log.Warnf("SyncPublishStart,NewSyncProducer error,err = %v", err)
		time.Sleep(3 * time.Second)
		go SyncPublishStart()
		return
	}
}

func AsyncPublishStart() {
	if KafkaCloseSign {
		return
	}
	var err error
	log.Info(base.GetConfig().Kafka.Hosts)
	*AsyncProducer, err = sarama.NewAsyncProducer(base.GetConfig().Kafka.Hosts, config())
	if err != nil {
		log.Warnf("AsyncProduce,NewSyncProducer error,err = %v", err)
		time.Sleep(3 * time.Second)
		go AsyncPublishStart()
		return
	}
}
