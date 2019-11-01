package img

import (
	"time"

	"github.com/gin-gonic/gin"
	"github.com/heyHui2018/log"

	"github.com/heyHui2018/best-practise/base"
	"github.com/heyHui2018/best-practise/model"

	imgM "github.com/heyHui2018/best-practise/model/img"
	imgS "github.com/heyHui2018/best-practise/service/img"
)

/*
param:file 图片文件
      x0 宽度起始位置,>0
      x1 宽度终止位置,>x0
      y0 高度起始位置,>0
      y1 高度终止位置,>y0
      quality 质量,取值0-100,越大越清晰
	  (左上角为(0,0)点)
*/

func Cut(c *gin.Context) {
	t := new(log.TLog)
	t.TraceId = c.GetString("traceId")
	start := time.Now()

	cut := new(imgM.Cut)
	err := c.ShouldBind(cut)
	if err != nil {
		t.Warnf("Cut ShouldBind error,err = %v", err)
		model.Fail(base.BadRequest, c)
		return
	}
	t.Infof("Cut 入参 cut = %+v", cut)

	m := new(imgM.Image)
	if !cut.Check(m) {
		t.Warnf("Cut params error")
		model.Fail(base.ParamError, c)
		return
	}

	file, fileHead, err := c.Request.FormFile("file")
	if err != nil {
		t.Warnf("Cut FormFile error,err = %v", err)
		model.Fail(base.ParamError, c)
		return
	}
	fileType := fileHead.Header.Get("Content-MarkType")
	t.Infof("Cut,fileType = %v", fileType) // fileType:img/jpeg,do some check here

	err = imgS.Decode(file, m)
	if err != nil {
		t.Warnf("Cut Decode error,err = %v", err)
		model.Fail(base.SystemError, c)
		return
	}

	out, err := imgS.Cut(t, m)
	if err != nil {
		model.Fail(base.SystemError, c)
		return
	}
	t.Infof("Cut 完成,耗时 = %v", time.Since(start))
	c.Writer.Write(out.Bytes())
}
