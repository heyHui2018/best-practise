package img

import (
	"bytes"
	"image"

	"github.com/disintegration/imaging"
	"github.com/heyHui2018/log"

	img "github.com/heyHui2018/best-practise/model/img"
)

func Cut(t *log.TLog, img *img.Image) (*bytes.Buffer, error) {
	rect := new(image.Rectangle)
	rect.Min.X = img.X0
	rect.Min.Y = img.Y0
	rect.Max.X = img.X1
	rect.Max.Y = img.Y1
	img.Data = imaging.Crop(img.Img, *rect)
	writer, err := Encode(img)
	if err != nil {
		t.Warnf("Cut Encode error,err = %v,format = %v", err, img.Format)
		return nil, err
	}
	return writer, nil
}
