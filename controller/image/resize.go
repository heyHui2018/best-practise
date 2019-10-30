package image

import (
	"time"

	"github.com/gin-gonic/gin"
	"github.com/heyHui2018/log"

	"github.com/heyHui2018/best-practise/base"
	"github.com/heyHui2018/best-practise/model"

	imgM "github.com/heyHui2018/best-practise/model/image"
	imgS "github.com/heyHui2018/best-practise/service/image"
)

/*
param:file 图片文件
      height 高度
      width 宽度
*/

func Resize(c *gin.Context) {
	t := new(log.TLog)
	t.TraceId = c.GetString("traceId")
	start := time.Now()
	// file, _, err := c.Request.FormFile("file")
	// if err != nil {
	// 	t.Warnf("Resize FormFile error,err = %v", err)
	// 	model.Fail(base.ParamError, c)
	// 	return
	// }
	r := new(imgM.Resize)
	err := c.ShouldBind(r)
	if err != nil {
		t.Warnf("Resize 入参 error,err = %v", err)
		model.Fail(base.BadRequest, c)
		return
	}
	r.File, _, err = c.Request.FormFile("file")
	r.Width = c.GetInt("width")
	r.Height = c.GetInt("height")
	t.Infof("123123 r = %+v", r)
	m := new(imgM.Image)

	// check fileHead
	// check md5

	m.Quality = 60
	m.Width = 300
	m.Height = 300
	err = imgS.Decode(r.File, m)
	if err != nil {
		t.Warnf("Resize Decode error,err = %v", err)
		model.Fail(base.SystemError, c)
		return
	}
	out := imgS.Resize(t, m)
	t.Infof("Resize 完成,耗时 = %v", time.Since(start))
	c.Writer.Write(out.Bytes())
	// model.Success(nil, c)
}
