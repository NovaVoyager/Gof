package middlewares

import (
	"bytes"
	"encoding/json"
	"github.com/gin-gonic/gin"
	"github.com/miaogu-go/Gof/global"
	"github.com/miaogu-go/Gof/utils"
	"go.uber.org/zap"
	"io/ioutil"
)

const (
	TraceIdKey = "trace_id"
)

//bodyLogWrite response结果存储
type bodyLogWrite struct {
	gin.ResponseWriter
	body *bytes.Buffer
}

//Write 实现写入方法
func (this bodyLogWrite) Write(b []byte) (int, error) {
	this.body.Write(b) //把响应写入自定义存储的地方保存
	return this.ResponseWriter.Write(b)
}

//RequestLog 请求日志
func RequestLog() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		body, err := ioutil.ReadAll(ctx.Request.Body)
		if err != nil {
			global.GofLog.Warn("get request body failed,err:", zap.Error(err))
		}
		ctx.Request.Body = ioutil.NopCloser(bytes.NewBuffer(body))

		//记录响应结果
		blw := &bodyLogWrite{ResponseWriter: ctx.Writer, body: bytes.NewBufferString("")}
		ctx.Writer = blw
		//生成traceId
		traceId := utils.GenerateTraceId()
		ctx.Set(TraceIdKey, traceId)
		ctx.Next()

		host := ctx.Request.Host
		query := ctx.Request.URL.RawQuery
		uri := ctx.Request.RequestURI
		method := ctx.Request.Method
		proto := ctx.Request.Proto
		ip := ctx.ClientIP()
		pid := utils.GetPid()
		path := ctx.Request.URL.Path

		bodyMap := make(map[string]interface{})
		if len(body) > 0 {
			err = json.Unmarshal(body, &bodyMap)
			if err != nil {
				global.GofLog.Warn("request body Unmarshal failed,err:", zap.Error(err))
			}
		}
		respMap := make(map[string]interface{})
		if blw.body.Len() > 0 {
			err = json.Unmarshal(blw.body.Bytes(), &respMap)
			if err != nil {
				global.GofLog.Warn("response body Unmarshal failed,err:", zap.Error(err))
			}
		}

		global.GofLog.Info("request log",
			zap.String("host", host),
			zap.String("uri", uri),
			zap.String("method", method),
			zap.String("query", query),
			zap.String("path", path),
			zap.String("proto", proto),
			zap.Any("head", ctx.Request.Header),
			zap.String("ip", ip),
			zap.String("trace_id", traceId),
			zap.Int("pid", pid),
			zap.Any("request_body", bodyMap),
			zap.Any("response_body", respMap),
		)
	}
}
