package main

import (
	"context"
	"github.com/gin-gonic/gin"
	"github.com/miaogu-go/Gof/core"
	"github.com/miaogu-go/Gof/global"
	"net/http"
	"time"
)

func main() {
	core.Bootstrap()
	r := gin.New()
	// 访问/version的返回值会随配置文件的变化而变化
	r.GET("/version", func(c *gin.Context) {
		c.String(http.StatusOK, global.GofConfig.System.Mode)
		global.GofRedis.Set(context.Background(), "a", "1", 30*time.Second)
		c.JSON(200, "")
	})

	err := r.Run(":8080")
	if err != nil {
		panic(err)
	}
}
