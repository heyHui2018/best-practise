package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/heyHui2018/utils"
	"github.com/ngaut/log"
	"net/http"
	"time"
)

func Weather(c *gin.Context) {
	traceId := c.GetString("traceId")
	start := time.Now()
	data := make(map[string]interface{})
	reply, err := utils.Get("https://http.cat/404", 20)
	if err != nil {
		log.Infof("Weather Get error,traceId = %v,err = %v", traceId, err)
		c.JSON(http.StatusInternalServerError, data)
	}
	log.Infof("Weather Get 完成,traceId = %v,reply = %v", traceId, string(reply)[:200])

	log.Infof("Weather 完成,traceId = %v,耗时 = %v", traceId, time.Since(start))
	c.JSON(http.StatusOK, reply)
}
