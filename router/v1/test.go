package v1

import (
	"github.com/gin-gonic/gin"
	v1Api "github.com/miaogu-go/Gof/api/v1"
)

func TestRoute(route *gin.RouterGroup) {
	testApi := v1Api.ApiV1App.TestGroup
	v1 := route.Group("v1")
	{
		v1.GET("/test", testApi.Add)
	}
}
