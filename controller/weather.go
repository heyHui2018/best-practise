package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/heyHui2018/best-practise/base"
	"github.com/heyHui2018/best-practise/models"
	"github.com/heyHui2018/utils"
	"github.com/ngaut/log"
	"time"
)

func Weather(c *gin.Context) {
	start := time.Now()
	traceId := c.GetString("traceId")
	city := c.Request.FormValue("city")
	state := c.Request.FormValue("state")
	country := c.Request.FormValue("country")
	log.Infof("Weather 入参,traceId = %v,param = %v", traceId, c.Request.Form)
	toCheck := map[string]string{
		"city":    city,
		"state":   state,
		"country": country,
	}
	if ok, param := utils.StrLengthCheck(toCheck); !ok {
		log.Warnf("Weather 入参 %v 为空,traceId = %v", param, traceId)
		models.Fail(base.MissingParam, c)
		return
	}
	avr := new(models.AirVisualReply)
	avr.Data.City = city
	avr.Data.State = state
	avr.Data.Country = country
	avd, err := avr.Data.Query()
	if err != nil {
		log.Warnf("Weather Query error,traceId = %v,err = %v", traceId, err)
		models.Fail(base.SystemError, c)
		return
	}
	log.Infof("Weather 完成,traceId = %v,avd = %+v,耗时 = %v", traceId, avd, time.Since(start))
	models.Success(avd, c)
}
