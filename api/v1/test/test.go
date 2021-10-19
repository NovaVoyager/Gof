package test

import (
	"github.com/gin-gonic/gin"
)

type Test struct {
	base
}

func (this *Test) Add(ctx *gin.Context) {
	//ctx.JSON(200, map[string]string{"a": "1"})
	//this.Ok(ctx)
	//this.OkWithMessage(ctx,"aaaaa")
	//this.OkWithData(ctx,map[string]string{"test":"this is a test request"})
	//this.OkWithData(ctx,S{Msg: "aaaaaaaaa"})
	this.OkWithDetailed(ctx, "res", map[string]interface{}{"key": 100})
	//this.Failed(ctx)
	//this.FailWithMessage(ctx,"aaaaa")
	//this.FailWithDetailed(ctx,"err",map[string]interface{}{"key":"value","age":100})
	//this.Fail401(ctx,"no auth")
}
