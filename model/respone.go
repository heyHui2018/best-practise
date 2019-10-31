package model

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"github.com/heyHui2018/best-practise/base"
)

type Response struct {
	Status int          `json:"status"`
	Msg    string       `json:"msg"`
	Data   interface{}  `json:"data"`
	G      *gin.Context `json:"-"`
}

func (r *Response) Res() {
	r.G.JSON(http.StatusOK, r)
}

func Success(data interface{}, g *gin.Context) {
	r := new(Response)
	r.Status = base.Success
	r.Msg = base.CodeText[base.Success]
	r.Data = data
	r.G = g
	r.Res()
}

func Fail(status int, g *gin.Context) {
	r := new(Response)
	r.Status = status
	r.Msg = base.CodeText[status]
	r.G = g
	r.Res()
}

func FailWithData(status int, data interface{}, g *gin.Context) {
	r := new(Response)
	r.Status = status
	r.Msg = base.CodeText[status]
	r.Data = data
	r.G = g
	r.Res()
}
