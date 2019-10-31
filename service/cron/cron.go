package cron

import (
	"time"

	"github.com/heyHui2018/log"
	"github.com/heyHui2018/utils"
	"github.com/robfig/cron"

	"github.com/heyHui2018/best-practise/service/dataSource"
	"github.com/heyHui2018/best-practise/service/mail"
)

func CronInit() {
	go Cron()
}

func Cron() {
	c := cron.New()
	err := c.AddFunc("@hourly", func() {
		// err := c.AddFunc("*/10 * * * * * ", func() {
		t := new(log.TLog)
		t.TraceId = time.Now().Format("20060102150405") + utils.GetRandomString()
		dataSource.GetWeather(t, "Shanghai", "Shanghai", "China")
	})
	if err != nil {
		log.Warnf("Cron c.AddFunc error,err = %v", err)
	}
	err = c.AddFunc("@hourly", func() {
		// err = c.AddFunc("*/10 * * * * * ", func() {
		t := new(log.TLog)
		t.TraceId = time.Now().Format("20060102150405") + utils.GetRandomString()
		mail.SendMail(t)
	})
	if err != nil {
		log.Warnf("Cron c.AddFunc error,err = %v", err)
	}
	c.Start()
}
