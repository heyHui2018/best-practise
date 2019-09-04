package routers

import (
	"github.com/gin-gonic/gin"
	"github.com/heyHui2018/best-practise/controller"
	"github.com/heyHui2018/best-practise/middleWare"
)

func InitRouter() *gin.Engine {
	g := gin.Default()
	g.Use(gin.Recovery())
	g.Use(middleWare.GenTraceId())

	pre := g.Group("/bestPractise")
	{
		jwtAll := middleWare.JWTInit(middleWare.AllUserAuthorizator)
		g.POST("/login", jwtAll.LoginHandler)
		token := pre.Group("/token")
		{
			token.GET("/refresh", jwtAll.RefreshHandler)
		}
		user := pre.Group("/user")
		{
			user.GET("/info", jwtAll.MiddlewareFunc())
		}

		jwtAdmin := middleWare.JWTInit(middleWare.AdminAuthorizator)
		admin := pre.Group("/admin")
		admin.Use(jwtAdmin.MiddlewareFunc())
		{

		}

		api := pre.Group("/Api")
		{
			api.POST("/register", controller.Register)
		}
	}
	return g
}
