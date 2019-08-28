package base

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

const (
	Success      = 200
	SystemError  = 500
	MissingParam = 1001
	ParamError   = 1002

	AirVisualDataLock = "AirVisualDataLock"
)

var codeText = map[int]string{
	Success:      "Success",
	SystemError:  "System error",
	MissingParam: "Missing param",
	ParamError:   "Param error",
}

func ResSuccess(data interface{}, c *gin.Context) {
	res := make(map[string]interface{})
	res["status"] = Success
	res["msg"] = codeText[Success]
	res["data"] = data
	c.JSON(http.StatusOK, res)
}

func ResFail(code int, c *gin.Context) {
	res := make(map[string]interface{})
	res["status"] = code
	res["msg"] = codeText[code]
	c.JSON(http.StatusOK, res)
}

func ResFailWithData(code int, data interface{}, c *gin.Context) {
	res := make(map[string]interface{})
	res["status"] = code
	res["msg"] = codeText[code]
	res["data"] = data
	c.JSON(http.StatusOK, res)
}
