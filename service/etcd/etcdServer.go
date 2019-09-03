package etcd

import (
	dis "github.com/heyHui2018/demo/etcd/discovery"
	"github.com/heyHui2018/log"
	"time"
)

func EtcdService() {
	serviceName := "s-test"
	serviceInfo := dis.ServiceInfo{IP: "172.16.16.114"}

	s, err := dis.NewService(serviceName, serviceInfo, []string{
		"http://172.16.16.114:2379",
	})
	if err != nil {
		log.Errorf("EtcdService error,err = %v", err)
	}

	go func() {
		time.Sleep(time.Second * 20)
		s.Stop()
	}()

	s.Start()
}
