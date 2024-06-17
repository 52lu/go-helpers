package logutil

import (
	"context"
	"testing"
)

func TestDemo(t *testing.T) {
	// 设置日志器
	SetLogger(LogConfig{
		Path:      "./logs",
		Level:     LogLevelDebug,
		FileName:  "app",
		OutFormat: OutFormatJson,
		LumberJackConf: LumberJackConfig{
			MaxSize:    1,
			MaxBackups: 30,
			MaxAge:     30,
			Compress:   false,
		},
	})
	ctx := context.Background()

	Debug(ctx, "debug测试", map[string]interface{}{"name": "张三"})
	Info(ctx, "info测试", map[string]interface{}{"name": "张三"})
}
