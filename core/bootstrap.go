package core

import (
	"github.com/gin-gonic/gin"
	"github.com/miaogu-go/Gof/core/initialize"
	"github.com/miaogu-go/Gof/global"
	"net/http"
	"time"
)

//Bootstrap 启动入口
func Bootstrap() {
	initialize.LoadConfig()
	initialize.LoadLog()
	initialize.LoadDb()
	initialize.LoadRedis()
	initialize.LoadMongodb()
	runServer()
}

//runServer 运行服务
func runServer() {
	gin.SetMode(global.GofConfig.System.Mode)
	ginEngine := initialize.RegisterRouter()
	/*s := initServer(global.GofConfig.System.Host, ginEngine)
	global.GofLog.Error(s.ListenAndServe().Error())*/
	global.GofLog.Error(ginEngine.Run(global.GofConfig.System.Host).Error())
}

type server interface {
	ListenAndServe() error
}

//initServer 初始化服务
func initServer(address string, router *gin.Engine) server {
	return &http.Server{
		Addr:           address,
		Handler:        router,
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}
}
