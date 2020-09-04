package router

import (
	"moneylogapi/handler/sd"
	"moneylogapi/router/middleware"
	"net/http"

	"github.com/gin-gonic/gin"
)

// Load 初始 路由
func Load(g *gin.Engine, mw ...gin.HandlerFunc) *gin.Engine {
	// 设置中间件
	//在处理某些请求时可能因为程序 bug 或者其他异常情况导致程序 panic，这时候为了不影响下一次请求的调用，需要通过 gin.Recovery()来恢复 API 服务器
	g.Use(gin.Recovery())
	g.Use(middleware.NoCach)
	g.Use(middleware.Options)
	g.Use(middleware.Secure)
	g.Use(mw...)

	// 404
	g.NoRoute(func(c *gin.Context) {
		c.String(http.StatusNotFound, "api 路由错误")
	})

	svcd := g.Group("/sd")
	{
		svcd.GET("health", sd.HealthCheck)
		svcd.GET("disk", sd.DiskCheck)
		svcd.GET("cpu", sd.CPUCheck)
		svcd.GET("ram", sd.RAMCheck)
		// svcd.GET("sysinfocheck", sd.SYSInfoCheck)
	}

	return g
}
