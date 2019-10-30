package image

import (
	"bytes"

	"github.com/disintegration/imaging"
	"github.com/heyHui2018/log"

	img "github.com/heyHui2018/best-practise/model/image"
)

func Resize(t *log.TLog, img *img.Image) *bytes.Buffer {
	img.Data = imaging.Resize(img.Img, img.Width, img.Height, imaging.Lanczos)
	writer, err := Encode(img)
	if err != nil {
		t.Warnf("Resize Encode error,err = %v,format = %v", err, img.Format)
		return nil
	}
	return writer
}
