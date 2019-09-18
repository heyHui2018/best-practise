package etcd

import (
	"context"
	"errors"
	"go.etcd.io/etcd/clientv3"
	"sync"
	"time"
)

type Client struct {
	Client      *clientv3.Client
	Key         string
	ServiceList map[string]string
	lock        sync.Mutex
}

func NewClient(endpoints []string) (*Client, error) {
	cli, err := clientv3.New(clientv3.Config{
		Endpoints:   endpoints,
		DialTimeout: 5 * time.Second,
	})
	if err != nil {
		return nil, err
	}
	return &Client{
		Client:      cli,
		ServiceList: make(map[string]string),
	}, nil
}

func (this *Client) GetValue(key string) ([]string, error) {
	var val []string
	resp, err := this.Client.Get(context.Background(), key, clientv3.WithPrefix())
	if err != nil {
		return nil, err
	}
	this.Key = key
	go this.watcher()
	if resp == nil || len(resp.Kvs) == 0 {
		return nil, errors.New("resp == nil || len(resp.Kvs) == 0")
	}
	for i := range resp.Kvs {
		if resp.Kvs[i] != nil && len(resp.Kvs[i].Value) != 0 {
			this.SetServiceList(string(resp.Kvs[i].Key), string(resp.Kvs[i].Value))
			val = append(val, string(resp.Kvs[i].Value))
		}
	}
	return val, nil
}

func (this *Client) watcher() {
	rch := this.Client.Watch(context.Background(), this.Key, clientv3.WithPrefix())
	for wresp := range rch {
		for _, ev := range wresp.Events {
			switch ev.Type {
			case clientv3.EventTypePut:
				this.SetServiceList(string(ev.Kv.Key), string(ev.Kv.Value))
			case clientv3.EventTypeDelete:
				this.DelServiceList(string(ev.Kv.Key))
			}
		}
	}
}

func (this *Client) SetServiceList(key, val string) {
	this.lock.Lock()
	defer this.lock.Unlock()
	this.ServiceList[key] = string(val)
}

func (this *Client) DelServiceList(key string) {
	this.lock.Lock()
	defer this.lock.Unlock()
	delete(this.ServiceList, key)
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
