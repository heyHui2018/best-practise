package dataSource

import (
	"encoding/json"
	"fmt"
	"github.com/heyHui2018/best-practise/models"
	"github.com/heyHui2018/best-practise/service/influxdb"
	"github.com/heyHui2018/utils"
	"github.com/ngaut/log"
	"time"
)

func GetWeather(traceId, city, state, country string) *models.AirVisualReply {
	avr := new(models.AirVisualReply)
	url := fmt.Sprintf("http://api.airvisual.com/v2/city?city=%s&state=%s&country=%s&key=cb87fa6b-4e8e-4dc3-9c3b-285ae71dc72a", city, state, country)
	reply, err := utils.Get(url, 20)
	if err != nil {
		log.Infof("GetWeather Get error,traceId = %v,err = %v", traceId, err)
		return avr
	}
	log.Infof("GetWeather utils.Get 完成,traceId = %v,reply = %v", traceId, string(reply))
	err = json.Unmarshal(reply, &avr)
	if err != nil {
		log.Warnf("GetWeather json.Unmarshal error,traceId = %v,err = %v", traceId, err)
		return avr
	}
	err = avr.Data.Insert()
	if err != nil {
		log.Warnf("GetWeather InsertAirVisualData error,traceId = %v,err = %v", traceId, err)
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
	influxdb.Insert("weather", tags, fields, time.Now(), traceId)
	log.Infof("GetWeather 完成,traceId = %v,avr = %+v", traceId, avr)
	return avr
}
