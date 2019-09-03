package etcd

import (
	"fmt"
	dis "github.com/heyHui2018/demo/etcd/discovery"
	"github.com/heyHui2018/log"
	"time"
)

func EtcdClient() {
	m, err := dis.NewMaster([]string{
		"http://172.16.16.114:2379",
	}, "services/")
	if err != nil {
		log.Errorf("EtcdClient error,err = %v", err)
	}
	for {
		for k, v := range m.Nodes {
			fmt.Printf("node:%s, ip=%s\n", k, v.Info.IP)
		}
		fmt.Printf("nodes num = %d\n", len(m.Nodes))
		time.Sleep(time.Second * 5)
	}
}
