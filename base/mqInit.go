package base

import (
	"github.com/ngaut/log"
	"github.com/streadway/amqp"
	"sync"
	"time"
)

var MQWait *sync.WaitGroup
var ConsumeConn *amqp.Connection
var ConsumeChannel *amqp.Channel
var PublishConn *amqp.Connection
var PublishChannel *amqp.Channel

func ConsumeConnect() {
	log.Info("ConsumeConnect 开始连接")
	var err error
	username := GetConfig().MQs["consume"].Username
	password := GetConfig().MQs["consume"].Password
	ip := GetConfig().MQs["consume"].Ip
	port := GetConfig().MQs["consume"].Port
	host := GetConfig().MQs["consume"].Host
	mqUrl := "amqp://" + username + ":" + password + "@" + ip + ":" + port + "/" + host
a:
	ConsumeConn, err = amqp.Dial(mqUrl)
	if err != nil {
		log.Warnf("ConsumeConnect 连接MQ出错，err = %v", err)
		time.Sleep(3 * time.Second)
		goto a
	}
	ConsumeChannel, err = ConsumeConn.Channel()
	if err != nil {
		log.Warnf("ConsumeConnect 打开channel出错，err = %v", err)
		ConsumeConn.Close()
		time.Sleep(3 * time.Second)
		goto a
	}
	log.Info("ConsumeConnect 连接完成")
}

func PublishConnect() {
	log.Info("PublishConnect 开始连接")
	var err error
	username := GetConfig().MQs["publish"].Username
	password := GetConfig().MQs["publish"].Password
	ip := GetConfig().MQs["publish"].Ip
	port := GetConfig().MQs["publish"].Port
	host := GetConfig().MQs["publish"].Host
	mqUrl := "amqp://" + username + ":" + password + "@" + ip + ":" + port + "/" + host
a:
	PublishConn, err = amqp.Dial(mqUrl)
	if err != nil {
		log.Warnf("PublishConnect 连接MQ出错，err = %v", err)
		time.Sleep(3 * time.Second)
		goto a
	}
	PublishChannel, err = PublishConn.Channel()
	if err != nil {
		log.Warnf("PublishConnect 打开channel出错，err = %v", err)
		PublishConn.Close()
		time.Sleep(3 * time.Second)
		goto a
	}
	log.Info("PublishConnect 连接完成")
}
