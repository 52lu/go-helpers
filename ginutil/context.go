package ginutil

import (
	"context"
	"fmt"
	"time"
)

const (
	GinContextTraceId         = "gin-context-TraceId"         // traceId
	GinContextBeginTimeMilli  = "gin-context-BeginTimeMilli"  // 开始时间(毫秒)
	GinContextClientIp        = "gin-context-ClientIp"        // 客户端ip
	GinContextClientUserAgent = "gin-context-ClientUserAgent" // 客户端操作系统
	GinContextRequestUrlPath  = "gin-context-RequestUrlPath"  // 请求接口路由
)

/*
* @Description: 获取traceId
* @Author: LiuQHui
* @Param ctx
* @Return string
* @Date 2024-06-12 17:34:37
 */
func GetTractId(ctx context.Context) string {
	value, ok := ctx.Value(GinContextTraceId).(string)
	if ok {
		return value
	}
	return ""
}

/*
* @Description: 获取开始时间
* @Author: LiuQHui
* @Param ctx
* @Return time.Time
* @Date 2024-06-12 17:36:03
 */
func GetBeginTime(ctx context.Context) time.Time {
	value, ok := ctx.Value(GinContextBeginTimeMilli).(int64)
	if ok {
		return time.UnixMilli(value)
	}
	return time.Time{}
}

/*
* @Description: 获取开始时间(毫秒)
* @Author: LiuQHui
* @Param ctx
* @Return int64
* @Date 2024-06-12 17:39:25
 */
func GetBeginTimeMilli(ctx context.Context) int64 {
	value, ok := ctx.Value(GinContextBeginTimeMilli).(int64)
	if ok {
		return value
	}
	return 0
}

/*
* @Description: 获取耗时
* @Author: LiuQHui
* @Param ctx
* @Return string
* @Date 2024-06-12 18:34:41
 */
func GetUseTime(ctx context.Context) string {
	// 计算耗时
	beginTime := GetBeginTimeMilli(ctx)
	if beginTime == 0 {
		return ""
	}
	useTimeInt64 := time.Now().UnixMilli() - beginTime
	useTime := time.Duration(useTimeInt64) * time.Millisecond
	return fmt.Sprintf(" %.3f", useTime.Seconds())
}

/*
* @Description: 客户端ip
* @Author: LiuQHui
* @Param ctx
* @Return string
* @Date 2024-06-12 17:36:57
 */
func GetClientIp(ctx context.Context) string {
	value, ok := ctx.Value(GinContextClientIp).(string)
	if ok {
		return value
	}
	return ""
}

/*
* @Description: 客户端操作系统
* @Author: LiuQHui
* @Param ctx
* @Return string
* @Date 2024-06-12 17:37:20
 */
func GetClientUserAgent(ctx context.Context) string {
	value, ok := ctx.Value(GinContextClientUserAgent).(string)
	if ok {
		return value
	}
	return ""
}

/*
* @Description: 请求接口路由
* @Author: LiuQHui
* @Param ctx
* @Return string
* @Date 2024-06-12 17:37:39
 */
func GetRequestUrl(ctx context.Context) string {
	value, ok := ctx.Value(GinContextRequestUrlPath).(string)
	if ok {
		return value
	}
	return ""
}
