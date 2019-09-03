package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/heyHui2018/best-practise/base"
	"github.com/heyHui2018/best-practise/models"
	"github.com/heyHui2018/log"
	"github.com/heyHui2018/utils"
	"time"
)

/*
param:city 城市
      state 省
      country 国家
      email 邮箱
      hour 定时发送时间,24小时制
*/

func Register(c *gin.Context) {
	// l := new(base.LogP)
	// l.TraceId = time.Now().Format("20060102150405") + utils.GetRandomString()
	// l.LogP("Register 完成,耗时 = %v", 123)
	start := time.Now()
	traceId := c.GetString("traceId")
	rr := new(models.RegisterRecord)
	rr.Email = c.Request.FormValue("email")
	rr.Hour = c.Request.FormValue("hour")
	rr.City = c.Request.FormValue("city")
	rr.State = c.Request.FormValue("state")
	rr.Country = c.Request.FormValue("country")
	log.Infof("Register 入参,traceId = %v,rr = %+v", traceId, rr)
	toCheck := map[string]string{
		"email":   rr.Email,
		"hour":    rr.Hour,
		"city":    rr.City,
		"state":   rr.State,
		"country": rr.Country,
	}
	if ok, param := utils.StrLengthCheck(toCheck); !ok {
		log.Warnf("Register 入参 %v 为空,traceId = %v", param, traceId)
		models.Fail(base.MissingParam, c)
		return
	}
	// 查询是否已注册
	getRes, err := rr.GetByEmail()
	if err != nil {
		log.Warnf("Register GetByEmail error,traceId = %v,err = %v", traceId, err)
		models.Fail(base.SystemError, c)
		return
	}
	if getRes.Id > 0 {
		log.Infof("Register email已存在,traceId = %v", traceId)
		if getRes.Hour != rr.Hour {
			err = rr.UpdateByEmail()
			log.Warnf("Register UpdateByEmail error,traceId = %v,err = %v", traceId, err)
		}
	} else {
		err = rr.Insert()
		if err != nil {
			log.Warnf("Register Insert error,traceId = %v,err = %v", traceId, err)
			models.Fail(base.SystemError, c)
			return
		}
	}
	log.Infof("Register 完成,traceId = %v,耗时 = %v", traceId, time.Since(start))
	models.Success(nil, c)
}
