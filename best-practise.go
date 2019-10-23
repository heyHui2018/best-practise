package main

import (
	"flag"
	"fmt"
	"github.com/heyHui2018/best-practise/base"
	"github.com/heyHui2018/best-practise/controller/rpc"
	"github.com/heyHui2018/best-practise/middleWare"
	"github.com/heyHui2018/best-practise/pb"
	"github.com/heyHui2018/best-practise/routers"
	"github.com/heyHui2018/best-practise/service/cron"
	"github.com/heyHui2018/best-practise/service/etcd"
	"github.com/heyHui2018/best-practise/service/rabbitMQ"
	"github.com/heyHui2018/best-practise/service/stop"
	"github.com/heyHui2018/log"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
	"net"
	"net/http"
	"os"
	"os/signal"
	"syscall"
)

func main() {
	flag.Parse() // 用于优雅重启

	base.ConfigInit()
	base.LogInit()
	base.DbInit()
	rabbitMQ.MQInit()
	etcd.EtcdInit()
	cron.CronInit()
	// nsq.NsqInit()
	// kafka.KafkaInit()

	g := routers.InitRouter()
	g.Use(middleWare.Cors())

	var err error

	// rpc
	rpcPort := fmt.Sprintf(":%d", base.GetConfig().Server.RpcPort)
	listen, err := net.Listen("tcp", rpcPort)
	if err != nil {
		log.Fatalf("failed to listen,err = %v", err)
	}
	s := grpc.NewServer()
	pb.RegisterGetServer(s, &rpc.Server{})
	pb.RegisterUserServer(s, &rpc.Server{})
	reflection.Register(s)
	go s.Serve(listen)
	log.Infof("rpc start listening on %v", rpcPort)

	// http
	httpPort := fmt.Sprintf(":%d", base.GetConfig().Server.HttpPort)
	var listener net.Listener

	if *stop.Graceful {
		f := os.NewFile(3, "")
		listener, err = net.FileListener(f)
	} else {
		listener, err = net.Listen("tcp", httpPort)
	}
	if err != nil {
		log.Fatalf("failed to listen,err = %v", err)
	}
	server := &http.Server{
		Addr:    httpPort,
		Handler: g,
	}
	go server.Serve(listener)
	log.Infof("http start listening on %v", httpPort)

	signs := make(chan os.Signal)
	signal.Notify(signs, syscall.SIGKILL, syscall.SIGTERM, syscall.SIGINT)
	for {
		select {
		case sign := <-signs:
			log.Infof("Receive signal: %v", sign)
			// 此处设置的sign配置可根据实际情况修改
			if sign == syscall.SIGKILL || sign == syscall.SIGTERM || sign == syscall.SIGINT {
				stop.Stop(server)
			} else {
				stop.Restart(server, listener)
			}
		}
	}
}
