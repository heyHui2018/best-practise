package main

import (
	"fmt"
	"github.com/heyHui2018/best-practise/base"
	"github.com/heyHui2018/best-practise/routers"
	"github.com/heyHui2018/best-practise/service"
	"github.com/ngaut/log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func main() {
	base.ConfigInit()
	base.LogInit()
	base.DbInit()
	service.MQStart()

	httpPort := fmt.Sprintf(":%d", base.GetConfig().Server.HttpPort)
	server := &http.Server{
		Addr:         httpPort,
		Handler:      routers.InitRouter(),
		ReadTimeout:  time.Duration(base.GetConfig().Server.ReadTimeout),
		WriteTimeout: time.Duration(base.GetConfig().Server.WriteTimeout),
	}
	log.Infof("start listening on %s", httpPort)
	go server.ListenAndServe()

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
	service.ConsumeRegularCloseSign = true
	log.Info("资源清理成功，退出")
	os.Exit(0)
}
