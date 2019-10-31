package etcd

import (
	"context"
	"errors"
	"sync"

	"go.etcd.io/etcd/clientv3"
)

type Client struct {
	Client      *clientv3.Client
	Key         string
	ServiceList map[string]string
	lock        sync.Mutex
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
