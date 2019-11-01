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
	  quality 	质量,取值0-100,越大越清晰
	  type		缩放类型 1-按尺寸缩放;2-按比例缩放
	  ratio		缩放百分比,须大于0 小于100为缩小
      height 	高度
      width 	宽度
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
	fileType := fileHead.Header.Get("Content-MarkType")
	t.Infof("Resize fileType = %v", fileType) // fileType:img/jpeg,do some check here

	err = imgS.Decode(file, m)
	if err != nil {
		t.Warnf("Resize Decode error,err = %v", err)
		model.Fail(base.SystemError, c)
		return
	}

	if m.ResizeType == 2 {
		m.Width = (m.Img.Bounds().Max.X * m.Ratio) / 100
		m.Height = (m.Img.Bounds().Max.Y * m.Ratio) / 100
	}

	out, err := imgS.Resize(t, m)
	if err != nil {
		model.Fail(base.SystemError, c)
		return
	}
	t.Infof("Resize 完成,耗时 = %v", time.Since(start))
	c.Writer.Write(out.Bytes())
}
