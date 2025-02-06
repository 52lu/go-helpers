package gormtype

type OrmConnectConfig struct {
	ConnMaxLifetime int64  `json:"conn_max_lifetime"` // 连接的最大生命周期
	Connection      string `json:"connection"`        // 连接
	ConsoleLog      string `json:"console_log"`       // 控制台日志
	Driver          string `json:"driver"`            // 数据库类型
	LogLevel        string `json:"log_level"`         // 日志级别
	MaxIdleConns    int64  `json:"max_idle_conns"`    // 最大空闲连接数
	MaxOpenConns    int64  `json:"max_open_conns"`    // 最大打开连接数
	SlowSqlTime     int    `json:"slow_sql_time"`     // 慢日志(毫秒)
}
