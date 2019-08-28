package models

import (
	"errors"
	"github.com/heyHui2018/best-practise/base"
	"github.com/heyHui2018/utils"
	"time"
)

type AirVisualReply struct {
	Status string        `json:"status"`
	Data   AirVisualData `json:"data"`
}

type AirVisualData struct {
	City     string   `xorm:"index 'city'"    json:"city"`
	State    string   `xorm:"index 'state'"   json:"state"`
	Country  string   `xorm:"index 'country'" json:"country"`
	Location Location `xorm:"extends"         json:"location"`
	Current  Current  `xorm:"extends"         json:"current"`
}

type Location struct {
	Type        string    `xorm:"'type'"        json:"type"`
	Coordinates []float64 `xorm:"'coordinates'" json:"coordinates"`
}

type Current struct {
	Weather   Weather   `xorm:"extends" json:"weather"`
	Pollution Pollution `xorm:"extends" json:"pollution"`
}

type Weather struct {
	Timestamp     time.Time `xorm:"DateTime 'w_timestamp'" json:"ts"`
	Temperature   int       `xorm:"'temperature'"          json:"tp"`
	Pressure      int       `xorm:"'pressure'"             json:"pr"`
	Humidity      int       `xorm:"'humidity'"             json:"hu"` // %
	WindSpeed     int       `xorm:"'wind_speed'"           json:"ws"` // m/s
	WindDirection int       `xorm:"'wind_direction'"       json:"wd"` // as an angle of 360° (N=0, E=90, S=180, W=270)
}

type Pollution struct {
	Timestamp time.Time `xorm:"DateTime 'p_timestamp'" json:"ts"`
	AQIUS     int       `xorm:"'aqi_us'"               json:"aqius"`  // AQI value based on US EPA standard
	MainUS    string    `xorm:"'main_us'"              json:"mainus"` // main pollutant for US AQI
	AQICN     int       `xorm:"'aqi_cn'"               json:"aqicn"`  // AQI value based on China MEP standard
	MainCN    string    `xorm:"'main_cn'"              json:"maincn"` // main pollutant for Chinese AQI
}

func (this *AirVisualData) Insert() error {
	// 分布式,取得锁的可入库
	conn, err := utils.GetRedisConnWithoutPool(base.GetConfig().Redis.Ip, base.GetConfig().Redis.Port, base.GetConfig().Redis.Password)
	err = utils.LockWithTimeout(conn, base.AirVisualDataLock, 20)
	if err != nil {
		return errors.New("未取得redis锁")
	}
	_, err = base.DBEngine.Insert(this)
	return err
}

func (this *AirVisualData) Query() (*AirVisualData, error) {
	ard := new(AirVisualData)
	_, err := base.DBEngine.Where("city = ? and state = ? and country = ?", this.City, this.State, this.Country).Desc("id").Get(ard)
	return ard, err
}

/*
CREATE TABLE `air_visual_data` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `city` varchar(20) NOT NULL COMMENT '城市',
  `state` varchar(20) NOT NULL COMMENT '省',
  `country` varchar(20) NOT NULL COMMENT '国家',
  `type` varchar(20) NOT NULL COMMENT '类型',
  `coordinates` text NOT NULL COMMENT '经纬度',
  `w_timestamp` datetime NOT NULL COMMENT '时间-天气',
  `temperature` int(5) NOT NULL COMMENT '温度',
  `pressure` int(5) DEFAULT NULL COMMENT '气压',
  `humidity` int(5) DEFAULT NULL COMMENT '湿度',
  `wind_speed` int(5) DEFAULT NULL COMMENT '风速',
  `wind_direction` int(5) DEFAULT NULL COMMENT '风向',
  `p_timestamp` datetime NOT NULL COMMENT '时间-污染',
  `aqi_us` int(5) DEFAULT NULL COMMENT '美标AQI',
  `main_us` varchar(20) DEFAULT NULL COMMENT '美标主要污染物',
  `aqi_cn` int(5) DEFAULT NULL COMMENT '中标AQI',
  `main_cn` varchar(20) DEFAULT NULL COMMENT '中标主要污染物',
  PRIMARY KEY (`id`),
  KEY `idx_city` (`city`),
  KEY `idx_state` (`state`),
  KEY `idx_country` (`country`)
) ENGINE=InnoDB AUTO_INCREMENT=4 DEFAULT CHARSET=utf8 COMMENT='天气';
*/
