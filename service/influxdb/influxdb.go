package influxdb

import (
	"fmt"
	"github.com/heyHui2018/best-practise/base"
	"github.com/heyHui2018/log"
	"github.com/influxdata/influxdb/client/v2"
	"time"
)

func connect(log *log.TLog) (client.Client, error) {
	config := new(client.HTTPConfig)
	config.Addr = fmt.Sprintf("http://%s:%s", base.GetConfig().InfluxDB.Ip, base.GetConfig().InfluxDB.Port)
	config.Username = base.GetConfig().InfluxDB.Username
	config.Password = base.GetConfig().InfluxDB.Password
	return client.NewHTTPClient(*config)
}

func Insert(measurement string, tags map[string]string, fields map[string]interface{}, t time.Time, log *log.TLog) {
	cli, err := connect(log)
	if err != nil {
		log.Warnf("InfluxDBInsert NewHTTPClient error,err = %v", err)
		return
	}
	defer cli.Close()

	bp, err := client.NewBatchPoints(client.BatchPointsConfig{
		Database:  base.GetConfig().InfluxDB.Database,
		Precision: "s",
	})
	if err != nil {
		log.Warnf("InfluxDBInsert NewBatchPoints error,err = %v", err)
		return
	}

	pt, err := client.NewPoint(measurement, tags, fields, t)
	if err != nil {
		log.Warnf("InfluxDBInsert NewPoint error,err = %v", err)
		return
	}
	bp.AddPoint(pt)

	if err := cli.Write(bp); err != nil {
		log.Warnf("InfluxDBInsert Write error,err = %v", err)
		return
	}
	// cmd := fmt.Sprintf("SELECT * FROM %s LIMIT %d", measurement, 10)
	// res, err := Query(cmd, log)
	// if err == nil {
	// 	log.Infof("res = %v", res)
	// }
	log.Infof("InfluxDBInsert 完成")
}

func Query(cmd string, log *log.TLog) ([]client.Result, error) {
	var res []client.Result
	cli, err := connect(log)
	if err != nil {
		log.Warnf("InfluxDBQuery NewHTTPClient error,err = %v", err)
		return res, err
	}
	defer cli.Close()

	q, err := cli.Query(client.Query{
		Command:  cmd,
		Database: base.GetConfig().InfluxDB.Database,
	})
	if err != nil {
		log.Warnf("InfluxDBQuery Query error,err = %v", err)
		return res, err
	}
	if q.Error() != nil {
		log.Warnf("InfluxDBQuery Query error,err = %v", q.Error())
		return res, err
	}
	res = q.Results
	log.Infof("InfluxDBQuery 完成")
	return res, nil
}
