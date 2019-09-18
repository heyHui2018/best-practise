package etcd

import (
	"context"
	"go.etcd.io/etcd/clientv3"
	"time"
)

type Service struct {
	client        *clientv3.Client
	lease         clientv3.Lease
	leaseResp     *clientv3.LeaseGrantResponse
	cancelFunc    func()
	keepAliveChan <-chan *clientv3.LeaseKeepAliveResponse
	key           string
}

func NewService(endpoints []string, ttl int64) (*Service, error) {
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

	ser := new(Service)
	ser.client = cli
	if err := ser.setLease(ttl); err != nil {
		return nil, err
	}
	return ser, nil
}

// 设置租约
func (this *Service) setLease(ttl int64) error {
	lease := clientv3.NewLease(this.client)

	// 设置租约时间
	leaseResp, err := lease.Grant(context.TODO(), ttl)
	if err != nil {
		return err
	}

	// 设置续租
	ctx, cancelFunc := context.WithCancel(context.TODO())
	leaseRespChan, err := lease.KeepAlive(ctx, leaseResp.ID)
	if err != nil {
		return err
	}

	this.lease = lease
	this.leaseResp = leaseResp
	this.cancelFunc = cancelFunc
	this.keepAliveChan = leaseRespChan
	return nil
}

// 监听 续租情况
// func (this *Service) ListenLeaseRespChan() {
// 	for {
// 		select {
// 		case leaseKeepResp := <-this.keepAliveChan:
// 			if leaseKeepResp == nil {
// 				fmt.Printf("已经关闭续租功能\n")
// 				return
// 			}
// 		}
// 	}
// }

// 通过租约 注册服务
func (this *Service) PutService(key, val string) error {
	kv := clientv3.NewKV(this.client)
	this.key = key
	_, err := kv.Put(context.TODO(), key, val, clientv3.WithLease(this.leaseResp.ID))
	return err
}

// 撤销租约
func (this *Service) RevokeLease() error {
	this.cancelFunc()
	// time.Sleep(2 * time.Second)
	_, err := this.lease.Revoke(context.TODO(), this.leaseResp.ID)
	return err
}
