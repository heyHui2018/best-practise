package etcd

import (
	"context"
	"go.etcd.io/etcd/clientv3"
)

type Service struct {
	Client        *clientv3.Client
	Lease         clientv3.Lease
	LeaseResp     *clientv3.LeaseGrantResponse
	CancelFunc    func()
	KeepAliveChan <-chan *clientv3.LeaseKeepAliveResponse
	Key           string
}

// 设置租约
func (this *Service) SetLease(ttl int64) error {
	lease := clientv3.NewLease(this.Client)

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

	this.Lease = lease
	this.LeaseResp = leaseResp
	this.CancelFunc = cancelFunc
	this.KeepAliveChan = leaseRespChan
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
	kv := clientv3.NewKV(this.Client)
	this.Key = key
	_, err := kv.Put(context.TODO(), key, val, clientv3.WithLease(this.LeaseResp.ID))
	return err
}

// 撤销租约
func (this *Service) RevokeLease() error {
	this.CancelFunc()
	// time.Sleep(2 * time.Second)
	_, err := this.Lease.Revoke(context.TODO(), this.LeaseResp.ID)
	return err
}
