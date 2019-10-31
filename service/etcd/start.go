package etcd

import (
	"time"

	"github.com/heyHui2018/log"

	"github.com/heyHui2018/best-practise/base"
	"github.com/heyHui2018/best-practise/model/etcd"
)

func EtcdInit() {
	go RegisterStart()
}

func RegisterStart() {
	log.Info("RegisterStart 开始")
	endpoints := base.GetConfig().Etcd.Endpoints
	if len(endpoints) == 0 {
		log.Fatalf("RegisterStart,endpoints error,endpoints = %v", endpoints)
	}
	keys := base.GetConfig().Etcd.Keys
	ip := base.GetConfig().Etcd.Ip
	ser, err := NewService(endpoints, 5)
	if err != nil {
		log.Warnf("RegisterStart,NewService error,err = %v", err)
		time.Sleep(10 * time.Second)
		go RegisterStart()
		return
	}
	for k := range keys {
		err = ser.PutService(keys[k], ip)
		if err != nil {
			log.Warnf("RegisterStart,PutService error,err = %v", err)
		}
	}
	log.Info("RegisterStart 完成")
	go Monitor(ser)
}

func Monitor(ser *etcd.Service) {
	for {
		leaseKeepResp := <-ser.KeepAliveChan
		if leaseKeepResp == nil {
			log.Warnf("心跳已停止")
			go RegisterStart()
			return
		}
	}
}
