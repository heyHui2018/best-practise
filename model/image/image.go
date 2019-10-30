package image

import (
	"image"
	"mime/multipart"
)

type Resize struct {
	File   multipart.File
	Width  int `form:"width"    binding:"required"`
	Height int `form:"height"   binding:"required"`
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
	Img     image.Image
	Data    *image.NRGBA
}
