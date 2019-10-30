package etcd

import (
	"github.com/heyHui2018/best-practise/model/etcd"
	"go.etcd.io/etcd/clientv3"
	"time"
)

func NewClient(endpoints []string) (*etcd.Client, error) {
	cli, err := clientv3.New(clientv3.Config{
		Endpoints:   endpoints,
		DialTimeout: 5 * time.Second,
	})
	if err != nil {
		return nil, err
	}
	return &etcd.Client{
		Client:      cli,
		ServiceList: make(map[string]string),
	}, nil
}

/*

package main

import (
	"github.com/heyHui2018/best-practise/service/etcd"
	"github.com/heyHui2018/log"
	"log"
	"time"
)

func main() {
	endpoints := []string{"http://172.16.16.114:2379"}
	key := "timestamp"
	cli, err := etcd.NewClient(endpoints)
	if err != nil {
		log.Fatal(err)
	}
	vals, err := cli.GetValue(key)
	if err != nil {
		log.Fatal(err)
	}
	// do something with vals
	for {
		// watch the vals' change
		time.Sleep(time.Second * 3)
	}
}

*/
