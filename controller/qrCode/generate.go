package qrCode

import (
	"github.com/gin-gonic/gin"
	"github.com/heyHui2018/best-practise/base"
	"github.com/heyHui2018/best-practise/model"
	"github.com/heyHui2018/log"
	"github.com/skip2/go-qrcode"
	"time"
)

/*
param:url 跳转url(带http(s)://前缀方可跳转)
*/

func Generate(c *gin.Context) {
	t := new(log.TLog)
	t.TraceId = c.GetString("traceId")
	start := time.Now()
	url := c.Request.FormValue("url")
	t.Infof("Generate 入参,url = %v", url)
	data, err := qrcode.Encode(url, qrcode.Medium, 256)
	if err != nil {
		t.Warnf("Generate qrcode.Encode error,err = %v", err)
		model.Fail(base.SystemError, c)
		return
	}
	t.Infof("Generate 完成,耗时 = %v", time.Since(start))
	_, err = c.Writer.Write(data)
	if err != nil {
		t.Warnf("Generate Writer.Write error,err = %v", err)
		model.Fail(base.SystemError, c)
	}
}
