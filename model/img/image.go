package img

import (
	"image"
)

type Resize struct {
	Type    int            `form:"type"    binding:"required"` // 缩放类型 1-按尺寸缩放;2-按比例缩放
	Quality int            `form:"quality" binding:"required"`
	Ratio   int            `form:"ratio"` // 百分比,须大于0.小于100为缩小
	Width   int            `form:"width"`
	Height  int            `form:"height"`
}

func (r *Resize) Check(m *Image) bool {
	if r.Quality <= 0 || r.Quality > 100 {
		return false
	}
	m.Quality = r.Quality

	if r.Type != 1 && r.Type != 2 {
		return false
	}
	m.ResizeType = r.Type

	if r.Type == 2 && r.Ratio <= 0 {
		return false
	}
	m.Ratio = r.Ratio

	if r.Type == 1 && r.Height <= 0 {
		return false
	}
	m.Height = r.Height

	if r.Type == 1 && r.Width <= 0 {
		return false
	}
	m.Width = r.Width

	return true
}

type Cut struct {
	X0      int `form:"x0"  binding:"required"`
	X1      int `form:"x1"  binding:"required"`
	Y0      int `form:"y0"  binding:"required"`
	Y1      int `form:"y1"  binding:"required"`
	Quality int `form:"quality"  binding:"required"`
}

func (c *Cut) Check(m *Image) bool {
	if c.X0 < 0 {
		return false
	}
	m.X0 = c.X0

	if c.X1 <= c.X0 {
		return false
	}
	m.X1 = c.X1

	if c.Y0 < 0 {
		return false
	}
	m.Y0 = c.Y0

	if c.Y1 <= c.Y0 {
		return false
	}
	m.Y1 = c.Y1

	if c.Quality <= 0 || c.Quality > 100 {
		return false
	}
	m.Quality = c.Quality

	return true
}

type Watermark struct {
	Quality int    `form:"quality"  binding:"required"`
	Type    int    `form:"type"     binding:"required"`
	Text    string `form:"text"`
}

func (w *Watermark) Check(m *Image) bool {
	if w.Quality <= 0 || w.Quality > 100 {
		return false
	}
	m.Quality = w.Quality

	if w.Type != 1 && w.Type != 2 {
		return false
	}
	m.MarkType = w.Type

	return true
}

type Image struct {
	Format     string
	Width      int
	Height     int
	Quality    int
	X0         int
	X1         int
	Y0         int
	Y1         int
	ResizeType int         // 缩放类型 1-按尺寸缩放;2-按比例缩放
	Ratio      int         // 缩放比例
	MarkType   int         // 水印类型 1-图片;2-文字
	Mark       image.Image // 水印图
	Text       string      // 水印文字
	Img        image.Image // 原图
	Data       *image.NRGBA
}
