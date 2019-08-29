package models

import (
	"github.com/heyHui2018/best-practise/base"
	"time"
)

type RegisterRecord struct {
	Id         int64
	Email      string    `xorm:"'email'"`
	Hour       string    `xorm:"'hour'"`
	City       string    `xorm:"'city'"`
	State      string    `xorm:"'state'"`
	Country    string    `xorm:"'country'"`
	CreateTime time.Time `xorm:"DateTime created 'create_time'"`
	UpdateTime time.Time `xorm:"DateTime updated 'update_time'"`
}

func (this *RegisterRecord) GetByEmail() (*RegisterRecord, error) {
	rr := new(RegisterRecord)
	_, err := base.DBEngine.Where("email = ?", this.Email).Get(rr)
	return rr, err
}

func (this *RegisterRecord) FindByHour() ([]*RegisterRecord, error) {
	rrList := make([]*RegisterRecord, 0)
	err := base.DBEngine.Where("hour = ?", this.Hour).Find(&rrList)
	return rrList, err
}

func (this *RegisterRecord) UpdateByEmail() error {
	_, err := base.DBEngine.Cols("hour").Where("email = ?", this.Email).Update(this)
	return err
}

func (this *RegisterRecord) Insert() error {
	_, err := base.DBEngine.Insert(this)
	return err
}

/*
CREATE TABLE `register_record` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `email` varchar(50) NOT NULL COMMENT '邮箱',
  `hour` varchar(10) NOT NULL COMMENT '定时发送时间,24小时制',
  `city` varchar(20) NOT NULL COMMENT '城市',
  `state` varchar(20) NOT NULL COMMENT '省',
  `country` varchar(20) NOT NULL COMMENT '国家',
  `create_time` datetime DEFAULT NULL COMMENT '创建时间',
  `update_time` datetime DEFAULT NULL COMMENT '更新时间',
  PRIMARY KEY (`id`),
  UNIQUE KEY `idx_email` (`email`) USING BTREE,
  KEY `idx_hour` (`hour`)
) ENGINE=InnoDB AUTO_INCREMENT=4 DEFAULT CHARSET=utf8 COMMENT='注册表';
*/
