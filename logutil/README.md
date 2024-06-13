### 1.初始化
```go
	// 设置日志器
	logutil.SetLogger(logutil.LogConfig{
		Path:           "./logs",
		Level:          logutil.LogLevelDebug,
		FileName:       "app",
		FileTimeFormat: time.DateOnly,
		OutFormat:      logutil.OutFormatJson,
		LumberJackConf: logutil.LumberJackConfig{
			MaxSize:    1,
			MaxBackups: 5,
			MaxAge:     5,
			Compress:   true,
		},
	})
```

### 2.使用
```shell
	logutil.Debug(c, "获取header信息", body)
```