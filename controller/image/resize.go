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
param:file 		图片文件
      height 	高度
      width 	宽度
      quality 	质量,取值0-100,越大越清晰
*/

func Resize(c *gin.Context) {
	t := new(log.TLog)
	t.TraceId = c.GetString("traceId")
	start := time.Now()
	r := new(imgM.Resize)
	err := c.ShouldBind(r)
	if err != nil {
		t.Warnf("Resize ShouldBind error,err = %v", err)
		model.Fail(base.BadRequest, c)
		return
	}
	r.Width = c.Request.FormValue("width")
	r.Height = c.Request.FormValue("height")
	r.Quality = c.Request.FormValue("quality")
	t.Infof("Resize 入参 r = %+v", r)

	m := new(imgM.Image)
	if !r.Check(m) {
		t.Warnf("Resize params error")
		model.Fail(base.ParamError, c)
		return
	}

	file, fileHead, err := c.Request.FormFile("file")
	if err != nil {
		t.Warnf("Resize FormFile error,err = %v", err)
		model.Fail(base.ParamError, c)
		return
	}
	fileType := fileHead.Header.Get("Content-Type")
	t.Infof("Resize fileType = %v", fileType)

	// fileType:such as image/jpeg,we can do some check here

	err = imgS.Decode(file, m)
	if err != nil {
		t.Warnf("Resize Decode error,err = %v", err)
		model.Fail(base.SystemError, c)
		return
	}
	out, err := imgS.Resize(t, m)
	if err != nil {
		model.Fail(base.SystemError, c)
		return
	}
	t.Infof("Resize 完成,耗时 = %v", time.Since(start))
	c.Writer.Write(out.Bytes())
}
