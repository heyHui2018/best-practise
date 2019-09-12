package main

import (
	"fmt"
	"github.com/heyHui2018/best-practise/base"
	"github.com/heyHui2018/best-practise/controller/rpc"
	"github.com/heyHui2018/best-practise/middleWare"
	"github.com/heyHui2018/best-practise/pb"
	"github.com/heyHui2018/best-practise/routers"
	"github.com/heyHui2018/best-practise/service/cron"
	"github.com/heyHui2018/best-practise/service/rabbitMQ"
	"github.com/heyHui2018/log"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
	"net"
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

	rpcPort := fmt.Sprintf(":%d", base.GetConfig().Server.RpcPort)
	listen, err := net.Listen("tcp", rpcPort)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	log.Infof("start listening on %v", rpcPort)
	s := grpc.NewServer()
	pb.RegisterGetServer(s, &rpc.Server{})
	reflection.Register(s)
	if err := s.Serve(listen); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}

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
