package routers

import (
	"github.com/gin-gonic/gin"

	"github.com/heyHui2018/best-practise/controller/dataSource"
	"github.com/heyHui2018/best-practise/controller/img"
	"github.com/heyHui2018/best-practise/controller/qrCode"
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
			api.POST("/register", dataSource.Register)
			api.POST("/generate", qrCode.Generate)

			m := api.Group("/img")
			{
				m.POST("/resize", img.Resize)
				m.POST("/cut", img.Cut)
				m.POST("/watermark", img.Watermark)
			}
		}
	}
	return g
}
