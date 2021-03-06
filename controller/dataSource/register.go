package dataSource

import (
	"time"

	"github.com/gin-gonic/gin"
	"github.com/heyHui2018/log"

	"github.com/heyHui2018/best-practise/base"
	"github.com/heyHui2018/best-practise/model"
	"github.com/heyHui2018/best-practise/model/dataSource"
)

/*
param:city 		城市
      state 	省
      country 	国家
      email 	邮箱
      hour 		定时发送时间,24小时制
*/

func Register(c *gin.Context) {
	t := new(log.TLog)
	t.TraceId = c.GetString("traceId")
	start := time.Now()
	rr := new(dataSource.RegisterRecord)
	/*
		bind：struct中添加 binding:"required"
		为了自定义返回,建议使用ShouldBind...
			表单：struct中添加 form:"xxx"
				post:ShouldBind
				get:ShouldBindQuery
			Json: struct中添加 json:"xxx"
				ShouldBindJSON
	*/
	err := c.ShouldBind(rr)
	if err != nil {
		t.Warnf("Register ShouldBind error,err = %v", err)
		model.Fail(base.BadRequest, c)
		return
	}
	t.Infof("Register 入参,rr = %+v", rr)

	// 查询是否已注册
	getRes, err := rr.GetByEmail()
	if err != nil {
		t.Warnf("Register GetByEmail error,err = %v", err)
		model.Fail(base.SystemError, c)
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
			model.Fail(base.SystemError, c)
			return
		}
	}
	t.Infof("Register 完成,耗时 = %v", time.Since(start))
	model.Success(nil, c)
}
