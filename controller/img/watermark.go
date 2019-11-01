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
param:file 		图片文件
      type 		水印类型 1-图片;2-文字
      quality 	质量,取值0-100,越大越清晰
	  mark		水印图 type=1时必须
	  text 		水印文字 type=2时必须
*/

func Watermark(c *gin.Context) {
	t := new(log.TLog)
	t.TraceId = c.GetString("traceId")
	start := time.Now()

	w := new(imgM.Watermark)
	err := c.ShouldBind(w)
	if err != nil {
		t.Warnf("Watermark ShouldBind error,err = %v", err)
		model.Fail(base.BadRequest, c)
		return
	}
	t.Infof("Watermark 入参 w = %+v", w)

	m := new(imgM.Image)
	if !w.Check(m) {
		t.Warnf("Watermark params error")
		model.Fail(base.ParamError, c)
		return
	}

	file, fileHead, err := c.Request.FormFile("file")
	if err != nil {
		t.Warnf("Watermark FormFile error,err = %v", err)
		model.Fail(base.ParamError, c)
		return
	}
	fileType := fileHead.Header.Get("Content-MarkType")
	t.Infof("Watermark fileType = %v", fileType) // fileType:such as img/jpeg,we can do some check here

	err = imgS.Decode(file, m)
	if err != nil {
		t.Warnf("Watermark Decode error,err = %v", err)
		model.Fail(base.SystemError, c)
		return
	}

	switch m.MarkType {
	case 1:
		mark, _, err := c.Request.FormFile("mark")
		if err != nil {
			t.Warnf("Watermark FormFile error,err = %v", err)
			model.Fail(base.ParamError, c)
			return
		}

		mm := new(imgM.Image)
		err = imgS.Decode(mark, mm)
		if err != nil {
			t.Warnf("Watermark Decode error,err = %v", err)
			model.Fail(base.SystemError, c)
			return
		}
		m.Mark = mm.Img
	case 2:
		if len(w.Text) == 0 {
			t.Warnf("Watermark w.Text 为空")
			model.Fail(base.ParamError, c)
			return
		}
		m.Text = w.Text
	default:
		t.Warnf("Watermark wrong m.MarkType,m.MarkType = %v", m.MarkType)
		model.Fail(base.ParamError, c)
		return
	}
	out, err := imgS.Watermark(t, m)
	if err != nil {
		model.Fail(base.SystemError, c)
		return
	}
	t.Infof("Watermark 完成,耗时 = %v", time.Since(start))
	c.Writer.Write(out.Bytes())
}
