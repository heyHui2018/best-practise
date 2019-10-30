package image

import (
	"bytes"
	"image"
	"image/color"
	"image/draw"
	"io/ioutil"

	"github.com/golang/freetype"
	"github.com/heyHui2018/log"

	img "github.com/heyHui2018/best-practise/model/image"
)

func Watermark(t *log.TLog, img *img.Image) (*bytes.Buffer, error) {
	b := img.Img.Bounds()
	m := image.NewNRGBA(b)
	switch img.Type {
	case 1:
		draw.Draw(m, b, img.Img, image.ZP, draw.Src)
		// image.ZP指(0,0)点 draw.Src指原图替换掉目标图
		draw.Draw(m, img.Mark.Bounds().Add(image.ZP), img.Mark, image.ZP, draw.Over) // draw.Over指原图覆盖在目标图上
	case 2:
		for y := 0; y < img.Img.Bounds().Dy(); y++ {
			for x := 0; x < img.Img.Bounds().Dx(); x++ {
				m.Set(x, y, img.Img.At(x, y))
			}
		}
		fontBytes, err := ioutil.ReadFile("conf/simsun.ttc")
		if err != nil {
			t.Warnf("Watermark ReadFile error,err = %v", err)
			return nil, err
		}
		font, err := freetype.ParseFont(fontBytes)
		if err != nil {
			t.Warnf("Watermark freetype.ParseFont error,err = %v", err)
			return nil, err
		}
		f := freetype.NewContext()
		f.SetDPI(72)
		f.SetFont(font)
		f.SetFontSize(20)
		f.SetClip(img.Img.Bounds())
		f.SetDst(m)
		f.SetSrc(image.NewUniform(color.RGBA{R: 255, G: 0, B: 0, A: 255}))
		pt := freetype.Pt(img.Img.Bounds().Dx()-200, img.Img.Bounds().Dy()-12)
		_, err = f.DrawString(img.Text, pt)
	}
	img.Data = m
	writer, err := Encode(img)
	if err != nil {
		t.Warnf("Watermark Encode error,err = %v,format = %v", err, img.Format)
		return nil, err
	}
	return writer, nil
}
