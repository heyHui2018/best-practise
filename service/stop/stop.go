package stop

import (
	"context"
	"flag"
	"github.com/heyHui2018/best-practise/service/kafka"
	"github.com/heyHui2018/best-practise/service/nsq"
	"github.com/heyHui2018/best-practise/service/rabbitMQ"
	"github.com/heyHui2018/log"
	"net"
	"net/http"
	"os"
	"os/exec"
	"reflect"
	"time"
)

/*
http服务优雅重启详见：https://github.com/heyHui2018/graceful
*/

var Graceful = flag.Bool("graceful", false, "listen on fd open 3 (internal use only)")

func StopInput() {
	// 关闭mq
	rabbitMQ.MQCloseSign = true
	if rabbitMQ.ConsumeChannel != nil {
		rabbitMQ.ConsumeChannel.Close()
	}
	if rabbitMQ.ConsumeConn != nil {
		rabbitMQ.ConsumeConn.Close()
	}
	// 关闭nsq
	nsq.NsqCloseSign = true
	if nsq.Consumer != nil {
		nsq.Consumer.Stop()
	}
	// 关闭kafka
	kafka.KafkaCloseSign = true
	if kafka.Consumer != nil {
		kafka.Consumer.Close()
	}
	wait := make(chan string)
	go func(wait chan string) {
		select {
		case <-wait:
			log.Info("资源清理成功，退出")
			break
		case <-time.After(3 * time.Second):
			log.Info("等待超时，退出")
		}
		StopOutput()
		os.Exit(0)
	}(wait)
	rabbitMQ.MQWait.Wait()
	nsq.NsqWait.Wait()
	kafka.KafkaWait.Wait()
	wait <- ""
}

func StopOutput() {
	if rabbitMQ.PublishChannel != nil {
		rabbitMQ.PublishChannel.Close()
	}
	if rabbitMQ.PublishConn != nil {
		rabbitMQ.PublishConn.Close()
	}
	if nsq.Producer != nil {
		nsq.Producer.Stop()
	}
	if *kafka.SyncProducer != nil {
		(*kafka.SyncProducer).Close()
	}
	if *kafka.AsyncProducer != nil {
		(*kafka.AsyncProducer).Close()
	}
}

func Stop(server *http.Server) {
	log.Info("开始停止程序...")
	// api优雅退出
	cxt, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	err := server.Shutdown(cxt)
	if err != nil {
		log.Warnf("Stop,server.Shutdown error,err = %v", err)
	}
	StopInput()
}

func Restart(server *http.Server, listener net.Listener) {
	log.Info("开始重启...")
	l, ok := listener.(*net.TCPListener)
	if !ok {
		log.Warnf("Restart,listener's type is not *net.TCPListener,type = %v", reflect.TypeOf(listener).Name())
		return
	}
	file, err := l.File()
	if err != nil {
		log.Warnf("Restart,l.File error,err = %v", err)
		return
	}
	args := os.Args
	if !*Graceful {
		args = append(args, "-graceful")
	}
	cmd := exec.Command(args[0], args[1:]...)
	cmd.Stdin = os.Stdin
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	cmd.ExtraFiles = []*os.File{file}
	err = cmd.Start()
	if err == nil {
		Stop(server)
	}
}
