package main

import (
	"fmt"
	"github.com/heyHui2018/best-practise/base"
	"github.com/heyHui2018/best-practise/middleWare"
	"github.com/heyHui2018/best-practise/routers"
	"github.com/heyHui2018/best-practise/service/cron"
	"github.com/heyHui2018/best-practise/service/rabbitMQ"
	"github.com/heyHui2018/log"
	"os"
	"os/signal"
	"syscall"
)

func main() {
	base.ConfigInit()
	base.LogInit()
	base.DbInit()
	rabbitMQ.MQInit()
	// base.DataInit()

	go cron.Cron()

	g := routers.InitRouter()
	g.Use(middleWare.Cors())

	httpPort := fmt.Sprintf(":%d", base.GetConfig().Server.HttpPort)
	go g.Run(httpPort)
	log.Infof("Start listening on %s", httpPort)

	signs := make(chan os.Signal)
	signal.Notify(signs, syscall.SIGKILL, syscall.SIGTERM, syscall.SIGHUP, os.Interrupt, os.Kill, os.Interrupt)
	for {
		msg := <-signs
		log.Infof("Receive signal: %v", msg)
		clear()
	}
}

func clear() {
	log.Info("开始停止程序....")
	rabbitMQ.ConsumeRegularCloseSign = true
	log.Info("资源清理成功，退出")
	os.Exit(0)
}
