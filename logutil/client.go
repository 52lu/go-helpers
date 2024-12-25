package logutil

import (
	"context"
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

/*
* @Description: 获取logger
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
func (l LoggerClient) writeMapContent(ctx context.Context, loglevel string, message string, content map[string]interface{}) {
	sugar := l.zapLoggerClient.zapLogger.Sugar()
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
