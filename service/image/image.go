package image

import (
	"bytes"
	"errors"
	"image"
	"image/gif"
	"image/jpeg"
	"image/png"
	"io"

	// _ "github.com/chai2010/webp"
	"golang.org/x/image/bmp"

	img "github.com/heyHui2018/best-practise/model/image"
)

func Decode(body io.Reader, img *img.Image) error {
	var err error
	img.Img, img.Format, err = image.Decode(body)
	return err
}

func Encode(img *img.Image) (*bytes.Buffer, error) {
	writer := bytes.NewBuffer(make([]byte, 0))
	var err error
	switch img.Format {
	case "jpeg":
		err = jpeg.Encode(writer, img.Data, &jpeg.Options{img.Quality})
	case "png":
		err = png.Encode(writer, img.Data)
	case "gif":
		err = gif.Encode(writer, img.Data, &gif.Options{})
	case "bmp":
		err = bmp.Encode(writer, img.Data)
	// case "webp":
	// 	err = webp.Encode(writer, img.Data, &webp.Options{})
	default:
		err = errors.New("wrong format")
	}
	return writer, err
}
