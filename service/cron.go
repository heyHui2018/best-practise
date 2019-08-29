package service

import (
	"github.com/heyHui2018/utils"
	"github.com/ngaut/log"
	"github.com/robfig/cron"
	"time"
)

func Cron() {
	c := cron.New()
	err := c.AddFunc("@hourly", func() {
		traceId := time.Now().Format("20060102150405") + utils.GetRandomString()
		GetWeather(traceId, "Shanghai", "Shanghai", "China")
	})
	if err != nil {
		log.Warnf("Cron c.AddFunc error,err = %v", err)
	}
	// err = c.AddFunc("@hourly", func() {
	err = c.AddFunc("*/10 * * * * * ", func() {
		traceId := time.Now().Format("20060102150405") + utils.GetRandomString()
		SendMail(traceId)
	})
	if err != nil {
		log.Warnf("Cron c.AddFunc error,err = %v", err)
	}
	c.Start()
}
