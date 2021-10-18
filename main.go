package main

import (
	"github.com/gin-gonic/gin"
	"github.com/miaogu-go/Gof/core"
)

func main() {
	core.Bootstrap()
	r := gin.New()
	// 访问/version的返回值会随配置文件的变化而变化
	r.GET("/version", func(c *gin.Context) {

	})

	err := r.Run(":8080")
	if err != nil {
		panic(err)
	}
}
