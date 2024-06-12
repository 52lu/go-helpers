package logutil

// 日志配置
type LogConfig struct {
	Path           string           `yaml:"path"`           // 日志文件目录
	Level          string           `yaml:"level"`          // 最低记录级别
	FilePrefix     string           `yaml:"filePrefix"`     // 日志文件前缀
	FileTimeFormat string           `yaml:"fileTimeFormat"` // 日志文件名中的日期格式；20060102
	OutFormat      string           `yaml:"outFormat"`      // 日志输出格式: json/console
	LumberJackConf LumberJackConfig `yaml:"lumberJackConf"` // 日志文件切割和压缩
}

// 日志切割
type LumberJackConfig struct {
	MaxSize    int  `yaml:"maxSize"`    // 单文件最大容量(单位MB)
	MaxBackups int  `yaml:"maxBackups"` // 保留旧文件的最大数量
	MaxAge     int  `yaml:"maxAge"`     // 旧文件最多保存几天
	Compress   bool `yaml:"compress"`   // 是否压缩/归档旧文件
}

// 日志输出格式
const (
	OutFormatJson    = "json"
	OutFormatConsole = "console"
)

// 日志级别
const (
	LogLevelDebug  = "debug"
	LogLevelInfo   = "info"
	LogLevelWarn   = "warn"
	LogLevelError  = "error"
	LogLevelDPanic = "dpanic"
	LogLevelPanic  = "panic"
	LogLevelFatal  = "fatal"
)
