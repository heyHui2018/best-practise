package img

import (
	"image"
	"strconv"
)

type Resize struct {
	Width   string `form:"width"    binding:"required"`
	Height  string `form:"height"   binding:"required"`
	Quality string `form:"quality"  binding:"required"`
}

func (r *Resize) Check(m *Image) bool {
	var err error
	m.Quality, err = strconv.Atoi(r.Quality)
	if err != nil || m.Quality <= 0 || m.Quality > 100 {
		return false
	}
	m.Height, err = strconv.Atoi(r.Height)
	if err != nil || m.Height <= 0 {
		return false
	}
	m.Width, err = strconv.Atoi(r.Width)
	if err != nil || m.Width <= 0 {
		return false
	}
	return true
}

type Cut struct {
	X0      string `form:"x0"  binding:"required"`
	X1      string `form:"x1"  binding:"required"`
	Y0      string `form:"y0"  binding:"required"`
	Y1      string `form:"y1"  binding:"required"`
	Quality string `form:"quality"  binding:"required"`
}

func (c *Cut) Check(m *Image) bool {
	var err error
	m.X0, err = strconv.Atoi(c.X0)
	if err != nil || m.X0 < 0 {
		return false
	}
	m.X1, err = strconv.Atoi(c.X1)
	if err != nil || m.X1 <= m.X0 {
		return false
	}
	m.Y0, err = strconv.Atoi(c.Y0)
	if err != nil || m.Y0 < 0 {
		return false
	}
	m.Y1, err = strconv.Atoi(c.Y1)
	if err != nil || m.Y1 <= m.Y0 {
		return false
	}
	m.Quality, err = strconv.Atoi(c.Quality)
	if err != nil || m.Quality <= 0 || m.Quality > 100 {
		return false
	}
	return true
}

type Watermark struct {
	Quality string `form:"quality"  binding:"required"`
	Type    string `form:"type"     binding:"required"`
	Text    string `form:"text"`
}

func (w *Watermark) Check(m *Image) bool {
	var err error
	m.Quality, err = strconv.Atoi(w.Quality)
	if err != nil || m.Quality <= 0 || m.Quality > 100 {
		return false
	}
	m.Type, err = strconv.Atoi(w.Type)
	if err != nil || (m.Type != 1 && m.Type != 2) {
		return false
	}
	return true
}

type Image struct {
	Format  string
	Width   int
	Height  int
	Quality int
	X0      int
	X1      int
	Y0      int
	Y1      int
	Img     image.Image // 原图
	Mark    image.Image // 水印图
	Text    string      // 水印文字
	Type    int         // 水印类型 1-图片;2-文字
	Data    *image.NRGBA
}
