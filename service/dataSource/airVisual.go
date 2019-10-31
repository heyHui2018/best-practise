package dataSource

import (
	"encoding/json"
	"fmt"
	"time"

	"github.com/heyHui2018/log"
	"github.com/heyHui2018/utils"

	"github.com/heyHui2018/best-practise/model/dataSource"
	"github.com/heyHui2018/best-practise/service/influxdb"
)

func GetWeather(t *log.TLog, city, state, country string) *dataSource.AirVisualReply {
	avr := new(dataSource.AirVisualReply)
	url := fmt.Sprintf("http://api.airvisual.com/v2/city?city=%s&state=%s&country=%s&key=cb87fa6b-4e8e-4dc3-9c3b-285ae71dc72a", city, state, country)
	reply, err := utils.Get(url, 20)
	if err != nil {
		t.Infof("GetWeather Get error,err = %v", err)
		return avr
	}
	t.Infof("GetWeather utils.Get 完成,reply = %v", string(reply))
	err = json.Unmarshal(reply, &avr)
	if err != nil {
		t.Warnf("GetWeather json.Unmarshal error,err = %v", err)
		return avr
	}
	err = avr.Data.Insert()
	if err != nil {
		t.Warnf("GetWeather InsertAirVisualData error,err = %v", err)
	}
	tags := make(map[string]string)
	tags["country"] = country
	tags["state"] = state
	tags["city"] = city
	fields := make(map[string]interface{})
	fields["timestamp"] = avr.Data.Current.Weather.Timestamp
	fields["temperature"] = avr.Data.Current.Weather.Temperature
	fields["pressure"] = avr.Data.Current.Weather.Pressure
	fields["humidity"] = avr.Data.Current.Weather.Humidity
	influxdb.Insert("weather", tags, fields, time.Now(), t)
	t.Infof("GetWeather 完成,avr = %+v", avr)
	return avr
}
