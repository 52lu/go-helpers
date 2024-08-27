package ginutil

import (
	"bytes"
	"github.com/52lu/go-helpers/ctxutil"
	"github.com/52lu/go-helpers/logutil"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"io"
	"strings"
	"time"
)

/*
* @Description: 追加附属信息到上下文
* @Author: LiuQHui
* @Param logger
* @Return gin.HandlerFunc
* @Date 2024-06-12 17:06:52
 */
func AdditionalMiddleware(ctx *gin.Context) {
	// 开始时间
	ctx.Set(ctxutil.GinContextBeginTimeMilli, time.Now().UnixMilli())
	// traceId
	ctx.Set(ctxutil.GinContextTraceId, strings.ReplaceAll(uuid.New().String(), "-", ""))
	// 客户端ip
	ctx.Set(ctxutil.GinContextClientIp, ctx.ClientIP())
	// RemoteIP
	ctx.Set(ctxutil.GinContextRemoteIp, ctx.RemoteIP())
	// 客户端操作系统
	ctx.Set(ctxutil.GinContextClientUserAgent, ctx.Request.UserAgent())
	// 请求接口路由
	ctx.Set(ctxutil.GinContextRequestUrlPath, ctx.Request.URL.Path)
	ctx.Next()
}

/*
* @Description: 请求入口记录
* @Author: LiuQHui
* @Param ctx
* @Date 2024-08-27 15:35:46
 */
func RequestEntranceLogMiddleware(ctx *gin.Context) {
	bodyBytes, _ := io.ReadAll(ctx.Request.Body)
	requestUri := ctx.Request.RequestURI
	logutil.Info(ctx, "RequestEntrance:"+requestUri, map[string]interface{}{
		"header":     ctx.Request.Header,
		"body":       string(bodyBytes),
		"remoteAddr": ctx.Request.RemoteAddr,
	})
	ctx.Request.Body = io.NopCloser(bytes.NewBuffer(bodyBytes))
	ctx.Next()
}
