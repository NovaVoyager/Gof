package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/miaogu-go/Gof/core"
	"github.com/miaogu-go/Gof/middlewares"
	"io/ioutil"
)

func main() {
	core.Bootstrap()
	r := gin.New()
	r.Use(middlewares.RequestLog())
	r.GET("/version", func(c *gin.Context) {
		body, err := ioutil.ReadAll(c.Request.Body)
		fmt.Println(string(body), err)
	})

	err := r.Run(":8080")
	if err != nil {
		panic(err)
	}
}
