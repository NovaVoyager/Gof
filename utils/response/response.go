package response

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

const (
	Error          = "-1"
	Success        = "0"
	DefaultOkMsg   = "ok"
	DefaultFailMsg = "error"
)

//ResponseData 响应返回数据
type ResponseData struct {
	Code    string      `json:"code"`
	Msg     string      `json:"msg"`
	Data    interface{} `json:"data"`
	TraceId string      `json:"trace_id"`
}

//Response 响应
type Response struct {
	resp     ResponseData
	httpCode int
}

func NewResponse(resp ResponseData, httpCode int) *Response {
	return &Response{resp: resp, httpCode: httpCode}
}

func (this *Response) defaultOkMsg(msg string) string {
	if msg == "" {
		return DefaultOkMsg
	}
	return msg
}

func (this *Response) defaultFailMsg(msg string) string {
	if msg == "" {
		return DefaultFailMsg
	}
	return msg
}

//result 返回响应
func (this *Response) result(ctx *gin.Context) {
	if this.httpCode == 0 {
		this.httpCode = http.StatusOK
	}
	if this.resp.Data == nil {
		this.resp.Data = struct{}{}
	}
	ctx.JSON(this.httpCode, this.resp)
}

//Ok 成功返回
func (this *Response) Ok(ctx *gin.Context) {
	this.resp.Code = Success
	this.resp.Msg = this.defaultOkMsg("")
	this.result(ctx)
}

//OkWithMessage 成功自定义消息返回
func (this *Response) OkWithMessage(ctx *gin.Context, msg string) {
	this.resp.Code = Success
	this.resp.Msg = this.defaultOkMsg(msg)
	this.result(ctx)
}

func (this *Response) OkWithData(ctx *gin.Context, data interface{}) {
	this.resp.Code = Success
	this.resp.Data = data
	this.resp.Msg = this.defaultOkMsg("")
	this.result(ctx)
}

func (this *Response) OkWithDetailed(ctx *gin.Context, msg string, data interface{}) {
	this.resp.Code = Success
	this.resp.Msg = this.defaultOkMsg(msg)
	this.resp.Data = data
	this.result(ctx)
}

func (this *Response) Failed(ctx *gin.Context) {
	this.resp.Code = Error
	this.resp.Msg = this.defaultFailMsg("")
	this.result(ctx)
}

func (this *Response) FailWithMessage(ctx *gin.Context, msg string) {
	this.resp.Code = Error
	this.resp.Msg = this.defaultFailMsg(msg)
	this.result(ctx)
}

func (this *Response) FailWithDetailed(ctx *gin.Context, msg string, data interface{}) {
	this.resp.Code = Error
	this.resp.Msg = this.defaultFailMsg(msg)
	this.resp.Data = data
	this.result(ctx)
}

func (this *Response) Fail401(ctx *gin.Context, msg string) {
	this.resp.Code = Error
	this.resp.Msg = this.defaultFailMsg(msg)
	this.httpCode = http.StatusUnauthorized
	this.result(ctx)
}

func (this *Response) Error(ctx *gin.Context, err error) {
	respErr, ok := err.(*RespError)
	if !ok {
		this.resp.Code = Error
		this.resp.Msg = this.defaultFailMsg("")
	} else {
		this.resp.Code = respErr.Code
		this.resp.Msg = this.defaultFailMsg(respErr.Msg)
	}
	this.result(ctx)
}
