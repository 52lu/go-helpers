package ginutil

import (
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
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
	ctx.Set(GinContextBeginTimeMilli, time.Now().UnixMilli())
	// traceId
	ctx.Set(GinContextTraceId, strings.ReplaceAll(uuid.New().String(), "-", ""))
	// 客户端ip
	ctx.Set(GinContextClientIp, ctx.ClientIP())
	// 客户端操作系统
	ctx.Set(GinContextClientUserAgent, ctx.Request.UserAgent())
	// 请求接口路由
	ctx.Set(GinContextRequestUrlPath, ctx.Request.URL.Path)
	ctx.Next()
}
