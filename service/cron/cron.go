package cron

import (
	"github.com/heyHui2018/best-practise/service/dataSource"
	"github.com/heyHui2018/best-practise/service/mail"
	"github.com/heyHui2018/log"
	"github.com/heyHui2018/utils"
	"github.com/robfig/cron"
	"time"
)

func Cron() {
	c := cron.New()
	err := c.AddFunc("@hourly", func() {
		traceId := time.Now().Format("20060102150405") + utils.GetRandomString()
		dataSource.GetWeather(traceId, "Shanghai", "Shanghai", "China")
	})
	if err != nil {
		log.Warnf("Cron c.AddFunc error,err = %v", err)
	}
	err = c.AddFunc("@hourly", func() {
		// err = c.AddFunc("*/10 * * * * * ", func() {
		traceId := time.Now().Format("20060102150405") + utils.GetRandomString()
		mail.SendMail(traceId)
	})
	if err != nil {
		log.Warnf("Cron c.AddFunc error,err = %v", err)
	}
	c.Start()
}
