package nsq

import (
	"sync"
	"time"

	"github.com/heyHui2018/log"
	"github.com/nsqio/go-nsq"

	"github.com/heyHui2018/best-practise/base"
)

var NsqCloseSign = false
var NsqWait sync.WaitGroup
var Producer *nsq.Producer
var Consumer *nsq.Consumer

func NsqInit() {
	go PublishStart()
	go ConsumeStart()
}

func PublishStart() {
	if NsqCloseSign {
		return
	}
	var err error
	Producer, err = nsq.NewProducer(base.GetConfig().Nsqs["publish"].Address, nsq.NewConfig())
	if err != nil {
		log.Warnf("NsqPublishStart,NewProducer error,err = %v", err)
		time.Sleep(3 * time.Second)
		go PublishStart()
		return
	}
	Producer.SetLogger(nil, 0) // 屏蔽系统日志
	err = Producer.Ping()
	if err != nil {
		log.Warnf("NsqPublishStart,Ping error,err = %v", err)
		Producer.Stop()
		Producer = nil
		time.Sleep(3 * time.Second)
		go PublishStart()
		return
	}
}

func ConsumeStart() {
	if NsqCloseSign {
		return
	}
	cfg := nsq.NewConfig()
	cfg.LookupdPollInterval = time.Second // 设置1s重连

	var err error
	Consumer, err = nsq.NewConsumer(base.GetConfig().Nsqs["consume"].Topic, base.GetConfig().Nsqs["consume"].Channel, cfg) // 新建一个消费者
	if err != nil {
		log.Warnf("NsqConsumeStart,NewConsumer error,err = %v", err)
		time.Sleep(3 * time.Second)
		NsqWait.Done()
		go ConsumeStart()
		return
	}
	Consumer.SetLogger(nil, 0) // 屏蔽系统日志

	Consumer.AddHandler(&ConsumerS{}) // 添加消费者接口

	Consumer.ChangeMaxInFlight(3) // 设置一次接收消息数量

	// 连接NSQLookupd
	err = Consumer.ConnectToNSQLookupd(base.GetConfig().Nsqs["consume"].Address)
	if err != nil {
		log.Warnf("NsqConsumeStart,ConnectToNSQLookupd error,err = %v", err)
		time.Sleep(3 * time.Second)
		NsqWait.Done()
		go ConsumeStart()
		return
	}
	// 连接多个nsqd
	// c.ConnectToNSQDs()
	// 连接单个nsqd
	// c.ConnectToNSQD()
}
