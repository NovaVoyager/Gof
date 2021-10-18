package main

import (
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
		type Jtw struct {
			Id        int
			CreatedAt time.Time
			UpdatedAt time.Time
			DeletedAt time.Time
			Jwt       string
		}
		jwt := new(Jtw)
		global.GofDB.Table("jwt_blacklists").First(jwt)
		c.JSON(200, jwt)
	})

	err := r.Run(":8080")
	if err != nil {
		panic(err)
	}
}
