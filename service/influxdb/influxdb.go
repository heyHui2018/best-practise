package influxdb

import (
	"fmt"
	"github.com/heyHui2018/best-practise/base"
	"github.com/heyHui2018/log"
	"github.com/influxdata/influxdb/client/v2"
	"time"
)

func Insert(measurement string, tags map[string]string, fields map[string]interface{}, t time.Time, tlog *log.TLog) {
	config := new(client.HTTPConfig)
	config.Addr = fmt.Sprintf("http://%s:%s", base.GetConfig().InfluxDB.Ip, base.GetConfig().InfluxDB.Port)
	config.Username = base.GetConfig().InfluxDB.Username
	config.Password = base.GetConfig().InfluxDB.Password

	cli, err := client.NewHTTPClient(*config)
	if err != nil {
		tlog.Warnf("InfluxDBInsert NewHTTPClient error,err = %v", err)
		return
	}
	defer cli.Close()

	bp, _ := client.NewBatchPoints(client.BatchPointsConfig{
		Database:  base.GetConfig().InfluxDB.Database,
		Precision: "s",
	})

	pt, err := client.NewPoint(measurement, tags, fields, t)
	if err != nil {
		tlog.Warnf("InfluxDBInsert NewPoint error,err = %v", err)
	}
	bp.AddPoint(pt)

	if err := cli.Write(bp); err != nil {
		tlog.Warnf("InfluxDBInsert Write error,err = %v", err)
	}
	tlog.Infof("InfluxDBInsert 完成")
}

func Query() {

}
