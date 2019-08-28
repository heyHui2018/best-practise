package middleWare

import (
	"github.com/gin-gonic/gin"
	"github.com/heyHui2018/utils"
	"net/http"
	"time"
)

// 跨域支持
func Cors() gin.HandlerFunc {
	return func(c *gin.Context) {
		method := c.Request.Method
		c.Header("Access-Control-Allow-Origin", "*")
		c.Header("Access-Control-Allow-Headers", "Content-Type,AccessToken,X-CSRF-Token,Authorization,Token,X-Token")
		c.Header("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE, UPDATE")
		c.Header("Access-Control-Expose-Headers", "Content-Length, Access-Control-Allow-Origin, Access-Control-Allow-Headers, Content-Type")
		c.Header("Access-Control-Allow-Credentials", "true")
		// 放行所有OPTIONS方法
		if method == "OPTIONS" {
			c.AbortWithStatus(http.StatusNoContent)
		}
		c.Next()
	}
}

// 生成追踪Id
func GenTraceId() gin.HandlerFunc {
	return func(c *gin.Context) {
		traceId := time.Now().Format("20060102150405") + utils.GetRandomString()
		c.Set("traceId", traceId)
		c.Next()
	}
}