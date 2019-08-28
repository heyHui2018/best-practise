package routers

import (
	"github.com/gin-gonic/gin"
	"github.com/heyHui2018/best-practise/base"
	"github.com/heyHui2018/best-practise/controller"
	"github.com/heyHui2018/best-practise/middleWare"
)

func InitRouter() *gin.Engine {
	r := gin.New()
	// r.Use(gin.Logger())
	r.Use(gin.Recovery())
	gin.SetMode(base.GetConfig().Server.RunMode)

	pre := r.Group("/bestPractise")
	pre.Use(middleWare.Cors())
	pre.Use(middleWare.GenTraceId())
	{
		bind := pre.Group("/Api")
		{
			bind.GET("/weather", controller.Weather)
		}
	}
	return r
}
