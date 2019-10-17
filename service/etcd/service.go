package etcd

import (
	"context"
	"github.com/heyHui2018/best-practise/models/etcd"
	"go.etcd.io/etcd/clientv3"
	"time"
)

func NewService(endpoints []string, ttl int64) (*etcd.Service, error) {
	cli, err := clientv3.New(clientv3.Config{
		Endpoints:   endpoints,
		DialTimeout: 3 * time.Second,
	})
	if err != nil {
		return nil, err
	}

	// 测试连接状态,因上一步clientv3.New即使在endpoints为空的情况下也不会报错(详见https://github.com/etcd-io/etcd/issues/9877),故在此需校验连接状态
	timeoutCtx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()
	_, err = cli.Status(timeoutCtx, endpoints[0])
	if err != nil {
		return nil, err
	}

	ser := new(etcd.Service)
	ser.Client = cli
	if err := ser.SetLease(ttl); err != nil {
		return nil, err
	}
	return ser, nil
}
