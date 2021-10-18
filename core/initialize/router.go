package initialize

import (
	"github.com/gin-gonic/gin"
	"github.com/miaogu-go/Gof/middlewares"
)

//RegisterRouter 注册路由
func RegisterRouter() *gin.Engine {
	r := gin.New()
	r.Use(middlewares.RequestLog(), middlewares.GinRecovery(true))

	publicGroup := r.Group("")
	{
		//健康监测
		publicGroup.GET("/health", func(ctx *gin.Context) {
			ctx.JSON(200, "ok")
		})
	}
	return r
}
