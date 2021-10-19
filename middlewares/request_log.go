package middlewares

import (
	"bytes"
	"encoding/json"
	"github.com/gin-gonic/gin"
	"github.com/miaogu-go/Gof/global"
	"go.uber.org/zap"
	"io/ioutil"
)

//RequestLog 请求日志
func RequestLog() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		body, err := ioutil.ReadAll(ctx.Request.Body)
		ctx.Request.Body = ioutil.NopCloser(bytes.NewBuffer(body))
		ctx.Next()

		host := ctx.Request.Host
		query := ctx.Request.URL.RawQuery
		uri := ctx.Request.RequestURI
		method := ctx.Request.Method
		proto := ctx.Request.Proto
		ip := ctx.ClientIP()
		path := ctx.Request.URL.Path
		if err != nil {
			global.GofLog.Warn("get request body failed,err:", zap.Error(err))
		}
		bodyMap := make(map[string]interface{})
		if len(body) > 0 {
			err = json.Unmarshal(body, &bodyMap)
			if err != nil {
				global.GofLog.Warn("body Unmarshal failed,err:", zap.Error(err))
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
			zap.Any("body", bodyMap),
		)
	}
}
