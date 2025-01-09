package logutil

import (
	"context"
	"fmt"
	"github.com/52lu/go-helpers/ctxutil"
)

var (
	_loggerClient *LoggerClient // 日志客户端
)

type LoggerClient struct {
	zapLoggerClient *zapLogClient
}

/*
* @Description: 设置日志器
* @Author: LiuQHui
* @Param cf
* @Date 2024-06-12 16:32:15
 */
func SetLogger(cf LogConfig) {
	zapClient, err := newZapLogClient(cf)
	if err != nil {
		return
	}
	_loggerClient = &LoggerClient{
		zapLoggerClient: zapClient,
	}
}

func GetLogger() *LoggerClient {
	return _loggerClient
}

/*
* @Description: 获取logger实例
* @Author: LiuQHui
* @Param cf
* @Return *LoggerClient
* @Return error
* @Date 2024-12-25 15:55:38
 */
func NewLogger(cf LogConfig) (*LoggerClient, error) {
	zapClient, err := newZapLogClient(cf)
	if err != nil {
		return nil, err
	}
	loggerClient := &LoggerClient{
		zapLoggerClient: zapClient,
	}
	return loggerClient, nil
}

func (l LoggerClient) Debug(ctx context.Context, message string, content map[string]interface{}) {
	l.writeMapContent(ctx, LogLevelDebug, message, content)
}

func (l LoggerClient) Debugf(ctx context.Context, message string, fmtArgs ...interface{}) {
	l.writeContentF(ctx, LogLevelDebug, message, fmtArgs...)
}

func (l LoggerClient) Info(ctx context.Context, message string, content map[string]interface{}) {
	l.writeMapContent(ctx, LogLevelInfo, message, content)
}

func (l LoggerClient) Infof(ctx context.Context, message string, fmtArgs ...interface{}) {
	l.writeContentF(ctx, LogLevelInfo, message, fmtArgs...)
}

func (l LoggerClient) Warn(ctx context.Context, message string, content map[string]interface{}) {
	l.writeMapContent(ctx, LogLevelWarn, message, content)
}

func (l LoggerClient) Warnf(ctx context.Context, message string, fmtArgs ...interface{}) {
	l.writeContentF(ctx, LogLevelWarn, message, fmtArgs...)
}

func (l LoggerClient) Error(ctx context.Context, message string, content map[string]interface{}) {
	l.writeMapContent(ctx, LogLevelError, message, content)
}

func (l LoggerClient) Errorf(ctx context.Context, message string, fmtArgs ...interface{}) {
	l.writeContentF(ctx, LogLevelError, message, fmtArgs...)
}

func (l LoggerClient) Printf(s string, fmtArgs ...interface{}) {
	l.writeContentF(context.Background(), LogLevelDebug, s, fmtArgs...)
}

/*
* @Description: 记录日志
* @Author: LiuQHui
* @Receiver l
* @Param ctx
* @Param loglevel
* @Param message
* @Param fmtArgs
* @Date 2024-06-12 18:00:29
 */
func (l LoggerClient) writeContentF(ctx context.Context, loglevel string, message string, fmtArgs ...interface{}) {
	content := fmt.Sprintf(message, fmtArgs...)
	l.writeMapContent(ctx, loglevel, content, nil)
}

/*
* @Description: 获取公共字段
* @Author: LiuQHui
* @Receiver l
* @Param ctx
* @Return map[string]interface{}
* @Date 2025-01-09 15:24:11
 */
func (l LoggerClient) getCommonField(ctx context.Context) map[string]interface{} {
	commonFieldMap := make(map[string]interface{})
	// traceId
	tractId := ctxutil.GetTractId(ctx)
	if tractId != "" {
		commonFieldMap["trace_id"] = tractId
	}
	// 请求耗时
	useTime := ctxutil.GetUseTime(ctx)
	if useTime != "" {
		commonFieldMap["use_time"] = useTime
	}
	// 客户端IP
	clientIp := ctxutil.GetClientIp(ctx)
	if clientIp != "" {
		commonFieldMap["client_ip"] = clientIp
	}
	// 客户端信息
	userAgent := ctxutil.GetClientUserAgent(ctx)
	if userAgent != "" {
		commonFieldMap["user_agent"] = userAgent
	}
	// 请求地址
	requestUrl := ctxutil.GetRequestUrl(ctx)
	if requestUrl != "" {
		commonFieldMap["request_url"] = requestUrl
	}
	return commonFieldMap
}

/*
* @Description: 记录map信息
* @Author: LiuQHui
* @Receiver l
* @Param ctx
* @Param loglevel
* @Param message
* @Param content
* @Date 2024-06-12 18:02:14
 */
func (l LoggerClient) writeMapContent(ctx context.Context, loglevel string, message string, content map[string]interface{}) {
	if content == nil {
		content = make(map[string]interface{})
	}
	////sugar := l.zapLoggerClient.zapLogger.Sugar()
	////defer sugar.Sync()
	//// 添加自定义字段
	//commonFieldMap := l.getCommonField(ctx)
	//for k, v := range commonFieldMap {
	//	content[k] = v
	//}
	zapField := getZapFieldCommonFromCtx(ctx)
	subLogger := l.zapLoggerClient.zapLogger.With(zapField...)
	sugar := subLogger.Sugar()
	defer sugar.Sync()

	key := "body"
	switch loglevel {
	case LogLevelDebug:
		sugar.Debugw(message, key, content)
	case LogLevelInfo:
		sugar.Infow(message, key, content)
	case LogLevelWarn:
		sugar.Warnw(message, key, content)
	case LogLevelError:
		sugar.Errorw(message, key, content)
	}
}
