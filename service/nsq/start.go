package nsq

import (
	"github.com/heyHui2018/best-practise/base"
	"github.com/heyHui2018/log"
	"github.com/nsqio/go-nsq"
	"sync"
	"time"
)

var NsqStopChan = make(chan string)
var NsqCloseSign = false
var NsqWait sync.WaitGroup
var Producer *nsq.Producer

func NsqInit() {
	go PublishStart()
	go ConsumeStart()
}

func PublishStart() {
	var err error
	Producer, err = nsq.NewProducer(base.GetConfig().Nsqs["publish"].Address, nsq.NewConfig())
	if err != nil {
		log.Warnf("PublishStart,NewProducer error,err = %v", err)
		time.Sleep(3 * time.Second)
		if !NsqCloseSign {
			go PublishStart()
		}
		return
	}
	Producer.SetLogger(nil, 0) // 屏蔽系统日志
	err = Producer.Ping()
	if err != nil {
		log.Warnf("PublishStart,Ping error,err = %v", err)
		Producer.Stop()
		Producer = nil
		time.Sleep(3 * time.Second)
		if !NsqCloseSign {
			go PublishStart()
		}
		return
	}
}

func ConsumeStart() {
	NsqWait.Add(1)
	cfg := nsq.NewConfig()
	cfg.LookupdPollInterval = time.Second // 设置1s重连

	c, err := nsq.NewConsumer(base.GetConfig().Nsqs["consume"].Topic, base.GetConfig().Nsqs["consume"].Channel, cfg) // 新建一个消费者
	if err != nil {
		log.Warnf("ConsumeStart,NewConsumer error,err = %v", err)
		time.Sleep(3 * time.Second)
		NsqWait.Done()
		if !NsqCloseSign {
			go ConsumeStart()
		}
		return
	}
	c.SetLogger(nil, 0) // 屏蔽系统日志

	c.AddHandler(&Consumer{}) // 添加消费者接口

	// 连接NSQLookupd

	err = c.ConnectToNSQLookupd(base.GetConfig().Nsqs["consume"].Address)
	if err != nil {
		log.Warnf("ConsumeStart,ConnectToNSQLookupd error,err = %v", err)
		time.Sleep(3 * time.Second)
		NsqWait.Done()
		if !NsqCloseSign {
			go ConsumeStart()
		}
		return
	}
	<-NsqStopChan
	c.Stop()
	NsqWait.Done()
	// 连接多个nsqd
	// c.ConnectToNSQDs()
	// 连接单个nsqd
	// c.ConnectToNSQD()
}
