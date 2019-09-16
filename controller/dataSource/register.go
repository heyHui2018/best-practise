package dataSource

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
	t := new(log.TLog)
	t.TraceId = c.GetString("traceId")
	start := time.Now()
	rr := new(models.RegisterRecord)
	rr.Email = c.Request.FormValue("email")
	rr.Hour = c.Request.FormValue("hour")
	rr.City = c.Request.FormValue("city")
	rr.State = c.Request.FormValue("state")
	rr.Country = c.Request.FormValue("country")
	t.Infof("Register 入参,rr = %+v", rr)
	toCheck := map[string]string{
		"email":   rr.Email,
		"hour":    rr.Hour,
		"city":    rr.City,
		"state":   rr.State,
		"country": rr.Country,
	}
	if ok, param := utils.StrLengthCheck(toCheck); !ok {
		t.Warnf("Register 入参 %v 为空", param)
		models.Fail(base.MissingParam, c)
		return
	}
	// 查询是否已注册
	getRes, err := rr.GetByEmail()
	if err != nil {
		t.Warnf("Register GetByEmail error,err = %v", err)
		models.Fail(base.SystemError, c)
		return
	}
	if getRes.Id > 0 {
		t.Infof("Register email已存在")
		if getRes.Hour != rr.Hour {
			err = rr.UpdateByEmail()
			t.Warnf("Register UpdateByEmail error,err = %v", err)
		}
	} else {
		err = rr.Insert()
		if err != nil {
			t.Warnf("Register Insert error,err = %v", err)
			models.Fail(base.SystemError, c)
			return
		}
	}
	t.Infof("Register 完成,耗时 = %v", time.Since(start))
	models.Success(nil, c)
}
