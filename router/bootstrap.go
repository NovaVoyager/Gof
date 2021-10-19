package router

import (
	"github.com/gin-gonic/gin"
	v1 "github.com/miaogu-go/Gof/router/v1"
)

//RouteBootstrap 路由加载入口文件
func RouteBootstrap(route *gin.RouterGroup) {
	publicGroup := route.Group("")
	{
		//健康监测
		publicGroup.GET("/health", func(ctx *gin.Context) {
			ctx.JSON(200, map[string]string{"msg": "ok"})
		})
	}

	v1.TestRoute(publicGroup)
}
