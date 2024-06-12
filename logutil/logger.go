package logutil

import (
	"context"
	"github.com/52lu/go-helpers/ginutil"
	"go.uber.org/zap"
	"sync"
)

var (
	_proxyLock = &sync.Mutex{}
)

/*
* @Description: 获取日志实例
* @Author: LiuQHui
* @Param ctx
* @Return *loggerClient
* @Date 2024-06-12 16:49:58
 */
func getLogger(ctx context.Context) *loggerClient {
	if _loggerClient == nil {
		_proxyLock.Lock()
		defer _proxyLock.Unlock()
		client, _ := newZapLogDefaultClient()
		_loggerClient = &loggerClient{
			zapLoggerClient: client,
		}
	}
	// 从上下文中获取信息
	_loggerClient.zapLoggerClient = addCommonFromCtx(ctx, _loggerClient.zapLoggerClient)
	return _loggerClient
}

/*
* @Description: 从上下文中获取公共信息
* @Author: LiuQHui
* @Param ctx
* @Param zapClient
* @Return *zapLogClient
* @Date 2024-06-12 17:33:30
 */
func addCommonFromCtx(ctx context.Context, zapClient *zapLogClient) *zapLogClient {
	zapLogger := zapClient.zapLogger
	var zapFields []zap.Field

	// traceId
	tractId := ginutil.GetTractId(ctx)
	if tractId != "" {
		zapFields = append(zapFields, zap.String("trace_id", tractId))
	}
	// 请求耗时
	useTime := ginutil.GetUseTime(ctx)
	if useTime != "" {
		zapFields = append(zapFields, zap.String("use_time", useTime))
	}
	// 客户端IP
	clientIp := ginutil.GetClientIp(ctx)
	if clientIp != "" {
		zapFields = append(zapFields, zap.String("client_ip", clientIp))
	}
	// 客户端信息
	userAgent := ginutil.GetClientUserAgent(ctx)
	if userAgent != "" {
		zapFields = append(zapFields, zap.String("user_agent", userAgent))
	}
	// 请求地址
	requestUrl := ginutil.GetRequestUrl(ctx)
	if requestUrl != "" {
		zapFields = append(zapFields, zap.String("request_url", requestUrl))
	}
	if len(zapFields) > 0 {
		zapLogger = zapLogger.With(zapFields...)
	}
	zapClient.zapLogger = zapLogger
	return zapClient
}

///*
//* @Description: 计算耗时
//* @Author: LiuQHui
//* @Param ctx
//* @Return string
//* @Date 2024-06-12 17:44:25
// */
//func computeUseTime(ctx context.Context) string {
//	// 计算耗时
//	beginTime := ginutil.GetBeginTimeMilli(ctx)
//	if beginTime == 0 {
//		return ""
//	}
//	useTimeInt64 := time.Now().UnixMilli() - beginTime
//	useTime := time.Duration(useTimeInt64) * time.Millisecond
//
//	return fmt.Sprintf(" %.3f", useTime.Seconds())
//}

/*
* @Description: Debug
* @Author: LiuQHui
* @Param ctx
* @Param message
* @Param content
* @Date 2024-06-12 16:46:42
 */
func Debug(ctx context.Context, message string, content map[string]interface{}) {
	getLogger(ctx).Debug(ctx, message, content)
}

/*
* @Description: Debugf
* @Author: LiuQHui
* @Param ctx
* @Param message
* @Param fmtArgs
* @Date 2024-06-12 16:47:23
 */
func Debugf(ctx context.Context, message string, fmtArgs ...interface{}) {
	getLogger(ctx).Debugf(ctx, message, fmtArgs...)
}

/*
* @Description: Info
* @Author: LiuQHui
* @Param ctx
* @Param message
* @Param content
* @Date 2024-06-12 16:47:59
 */
func Info(ctx context.Context, message string, content map[string]interface{}) {
	getLogger(ctx).Info(ctx, message, content)
}

/*
* @Description: Infof
* @Author: LiuQHui
* @Param ctx
* @Param message
* @Param fmtArgs
* @Date 2024-06-12 16:48:24
 */
func Infof(ctx context.Context, message string, fmtArgs ...interface{}) {
	getLogger(ctx).Infof(ctx, message, fmtArgs...)
}

/*
* @Description: Warn
* @Author: LiuQHui
* @Param ctx
* @Param message
* @Param content
* @Date 2024-06-12 16:49:34
 */
func Warn(ctx context.Context, message string, content map[string]interface{}) {
	getLogger(ctx).Warn(ctx, message, content)
}

/*
* @Description: Warnf
* @Author: LiuQHui
* @Param ctx
* @Param message
* @Param fmtArgs
* @Date 2024-06-12 16:49:25
 */
func Warnf(ctx context.Context, message string, fmtArgs ...interface{}) {
	getLogger(ctx).Warnf(ctx, message, fmtArgs...)
}

/*
* @Description: Error
* @Author: LiuQHui
* @Param ctx
* @Param message
* @Param content
* @Date 2024-06-12 16:49:28
 */
func Error(ctx context.Context, message string, content map[string]interface{}) {
	getLogger(ctx).Error(ctx, message, content)

}

/*
* @Description: Errorf
* @Author: LiuQHui
* @Param ctx
* @Param message
* @Param fmtArgs
* @Date 2024-06-12 16:49:27
 */
func Errorf(ctx context.Context, message string, fmtArgs ...interface{}) {
	getLogger(ctx).Errorf(ctx, message, fmtArgs...)
}
