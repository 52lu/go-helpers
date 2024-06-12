package logutil

import (
	"context"
	"github.com/52lu/go-helpers/jsonutil"
)

var (
	_loggerClient *loggerClient // 日志客户端
)

type loggerClient struct {
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
	_loggerClient = &loggerClient{
		zapLoggerClient: zapClient,
	}
}

func (l loggerClient) Debug(ctx context.Context, message string, content map[string]interface{}) {
	l.writeMapContent(ctx, LogLevelDebug, message, content)
}

func (l loggerClient) Debugf(ctx context.Context, message string, fmtArgs ...interface{}) {
	l.writeContentF(ctx, LogLevelDebug, message, fmtArgs...)
}

func (l loggerClient) Info(ctx context.Context, message string, content map[string]interface{}) {
	l.writeMapContent(ctx, LogLevelInfo, message, content)
}

func (l loggerClient) Infof(ctx context.Context, message string, fmtArgs ...interface{}) {
	l.writeContentF(ctx, LogLevelInfo, message, fmtArgs...)
}

func (l loggerClient) Warn(ctx context.Context, message string, content map[string]interface{}) {
	l.writeMapContent(ctx, LogLevelWarn, message, content)
}

func (l loggerClient) Warnf(ctx context.Context, message string, fmtArgs ...interface{}) {
	l.writeContentF(ctx, LogLevelWarn, message, fmtArgs...)
}

func (l loggerClient) Error(ctx context.Context, message string, content map[string]interface{}) {
	l.writeMapContent(ctx, LogLevelError, message, content)
}

func (l loggerClient) Errorf(ctx context.Context, message string, fmtArgs ...interface{}) {
	l.writeContentF(ctx, LogLevelError, message, fmtArgs...)
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
func (l loggerClient) writeContentF(ctx context.Context, loglevel string, message string, fmtArgs ...interface{}) {
	sugar := l.zapLoggerClient.zapLogger.Sugar()
	defer sugar.Sync()
	switch loglevel {
	case LogLevelDebug:
		sugar.Debugf(message, fmtArgs...)
	case LogLevelInfo:
		sugar.Infof(message, fmtArgs...)
	case LogLevelWarn:
		sugar.Warnf(message, fmtArgs...)
	case LogLevelError:
		sugar.Errorf(message, fmtArgs...)
	}
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
func (l loggerClient) writeMapContent(ctx context.Context, loglevel string, message string, content map[string]interface{}) {
	sugar := l.zapLoggerClient.zapLogger.Sugar()
	defer sugar.Sync()
	tmpContent, _ := jsonutil.Json.MarshalToString(content)
	switch loglevel {
	case LogLevelDebug:
		sugar.Debug(message, tmpContent)
	case LogLevelInfo:
		sugar.Info(message, tmpContent)
	case LogLevelWarn:
		sugar.Warn(message, tmpContent)
	case LogLevelError:
		sugar.Error(message, tmpContent)
	}
}
