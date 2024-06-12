package logutil

import (
	"github.com/52lu/go-helpers/pathutil"
	"github.com/natefinch/lumberjack"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"path"
	"strings"
	"time"
)

type zapLogClient struct {
	conf      LogConfig
	zapLogger *zap.Logger
}

/*
* @Description: 默认Zap实例
* @Author: LiuQHui
* @Return *zapLogClient
* @Return error
* @Date 2024-06-12 16:41:01
 */
func newZapLogDefaultClient() (*zapLogClient, error) {
	logger, _ := zap.NewProduction()
	return &zapLogClient{
		zapLogger: logger,
	}, nil
}

/*
* @Description: 实例化Zap实例
* @Author: LiuQHui
* @Param cf
* @Return *zapLogClient
* @Return error
* @Date 2024-06-12 15:42:54
 */
func newZapLogClient(cf LogConfig) (*zapLogClient, error) {
	client := &zapLogClient{
		conf: cf,
	}
	// 判断日志目录是否存在
	err := pathutil.CreateDir(cf.Path)
	if err != nil {
		return nil, err
	}
	// 设置输出格式
	var encoder zapcore.Encoder
	if cf.OutFormat == OutFormatJson {
		encoder = zapcore.NewJSONEncoder(client.getEncoderConfig())
	} else {
		encoder = zapcore.NewConsoleEncoder(client.getEncoderConfig())
	}
	// 设置日志文件切割
	writeSyncer := zapcore.AddSync(client.getLumberjackWriteSyncer())
	// 创建NewCore
	zapCore := zapcore.NewCore(encoder, writeSyncer, client.getLevel())
	// 创建logger
	client.zapLogger = zap.New(zapCore, zap.AddCaller(), zap.AddStacktrace(zap.ErrorLevel))

	return client, nil
}

/*
* @Description: 获取最低记录日志级别
* @Author: LiuQHui
* @Receiver z
* @Return zapcore.Level
* @Date 2024-06-12 15:50:30
 */
func (z *zapLogClient) getLevel() zapcore.Level {
	levelMap := map[string]zapcore.Level{
		LogLevelDebug:  zapcore.DebugLevel,
		LogLevelInfo:   zapcore.InfoLevel,
		LogLevelWarn:   zapcore.WarnLevel,
		LogLevelError:  zapcore.ErrorLevel,
		LogLevelDPanic: zapcore.DPanicLevel,
		LogLevelPanic:  zapcore.PanicLevel,
		LogLevelFatal:  zapcore.FatalLevel,
	}
	if level, ok := levelMap[z.conf.Level]; ok {
		return level
	}
	// 默认info级别
	return zapcore.InfoLevel
}

/*
* @Description: 自定义日志输出字段
* @Author: LiuQHui
* @Param cf
* @Return zapcore.EncoderConfig
* @Date 2024-06-12 14:19:54
 */
func (z *zapLogClient) getEncoderConfig() zapcore.EncoderConfig {
	config := zapcore.EncoderConfig{
		MessageKey:     "msg",
		LevelKey:       "level",
		TimeKey:        "time",
		NameKey:        "logger",
		CallerKey:      "caller",
		FunctionKey:    zapcore.OmitKey,
		StacktraceKey:  "stacktrace",
		LineEnding:     zapcore.DefaultLineEnding,
		EncodeLevel:    zapcore.CapitalLevelEncoder,
		EncodeTime:     z.getEncodeTime, // 自定义输出时间格式
		EncodeDuration: zapcore.StringDurationEncoder,
		EncodeCaller:   zapcore.ShortCallerEncoder,
	}
	return config
}

/*
* @Description: 定义日志输出时间格式
* @Author: LiuQHui
* @Param t
* @Param enc
* @Date 2024-06-12 14:19:27
 */
func (z *zapLogClient) getEncodeTime(t time.Time, enc zapcore.PrimitiveArrayEncoder) {
	enc.AppendString(t.Format("2006-01-02 15:04:05.000"))
}

/*
* @Description: 获取文件切割和归档配置信息
* @Author: LiuQHui
* @Param cf
* @Return zapcore.WriteSyncer
* @Date 2024-06-12 14:19:19
 */
func (z *zapLogClient) getLumberjackWriteSyncer() zapcore.WriteSyncer {
	lumberjackConfig := z.conf.LumberJackConf
	lumberjackLogger := &lumberjack.Logger{
		Filename:   z.getLogFile(),              //日志文件
		MaxSize:    lumberjackConfig.MaxSize,    //单文件最大容量(单位MB)
		MaxBackups: lumberjackConfig.MaxBackups, //保留旧文件的最大数量
		MaxAge:     lumberjackConfig.MaxAge,     // 旧文件最多保存几天
		Compress:   lumberjackConfig.Compress,   // 是否压缩/归档旧文件
	}
	// 设置日志文件切割
	return zapcore.AddSync(lumberjackLogger)
}

/*
* @Description: 获取日志文件名
* @Author: LiuQHui
* @Param cf
* @Return string
* @Date 2024-06-12 14:19:45
 */
func (z *zapLogClient) getLogFile() string {
	fileFormat := time.Now().Format(z.conf.FileTimeFormat)
	fileName := strings.Join([]string{
		z.conf.FilePrefix,
		fileFormat,
		"log"}, ".")
	return path.Join(z.conf.Path, fileName)
}
