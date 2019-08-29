package routers

import (
	"github.com/gin-gonic/gin"
	"github.com/heyHui2018/best-practise/controller"
	"github.com/heyHui2018/best-practise/middleWare"
)

func InitRouter() *gin.Engine {
	g := gin.Default()
	g.Use(gin.Recovery())
	pre := g.Group("/bestPractise")
	pre.Use(middleWare.GenTraceId())
	{
		api := pre.Group("/Api")
		{
			api.POST("/register", controller.Register)
		}
	}
	return g
}
