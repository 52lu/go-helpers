package logutil

import (
	"context"
)

type LoggerInterface interface {
	Debug(ctx context.Context, message string, context map[string]interface{})

	Debugf(ctx context.Context, message string, fmtArgs ...interface{})

	Info(ctx context.Context, message string, context map[string]interface{})

	Infof(ctx context.Context, message string, fmtArgs ...interface{})

	Warn(ctx context.Context, message string, context map[string]interface{})

	Warnf(ctx context.Context, message string, fmtArgs ...interface{})

	Error(ctx context.Context, message string, context map[string]interface{})

	Errorf(ctx context.Context, message string, fmtArgs ...interface{})
}
