package conftype

import "time"

// 接口
type ConfigParseInterface interface {
	Parse() error
	ConfigGetInterface
}

type ConfigGetInterface interface {
	Get(key string) interface{}
	GetString(key string) string
	GetBool(key string) bool
	GetInt64(key string) int64
	GetInt(key string) int
	GetFloat64(key string) float64
	GetTime(key string) time.Time
	GetIntSlice(key string) []int
	GetStringSlice(key string) []string
	GetStringMap(key string) map[string]interface{}
	GetStringMapString(key string) map[string]string
	GetStringMapStringSlice(key string) map[string][]string
}

// 配置
type ConfigParseConf struct {
	ConfigPaths []string        // 配置目录，优先级根据顺序来
	ConfigFile  string          // 配置文件
	ParseMethod ParseMethodType // 解析方式类型类型
}

// Apollo配置
type ApolloConfig struct {
	Enable            bool     `json:"enable" toml:"enable" yaml:"enable"`                                    // 是否启用Apollo配置
	ServiceUrl        string   `json:"service_url" toml:"serviceUrl" yaml:"serviceUrl"`                       // apollo服务地址
	Cluster           string   `json:"cluster" toml:"cluster" yaml:"cluster"`                                 // 集群
	AppId             string   `json:"app_id" toml:"appId" yaml:"appId"`                                      // appId
	Secret            string   `json:"secret" toml:"secret" yaml:"secret"`                                    // 安全模式下客户端需要的访问密钥
	SyncServerTimeout int      `json:"sync_server_timeout" toml:"syncServerTimeout" yaml:"syncServerTimeout"` // 同步配置服务的超时时间,默认为10秒
	Namespaces        []string `json:"namespaces" toml:"namespaces" yaml:"namespaces"`                        // 命名空间
	IsBackupConfig    bool     `json:"is_backup_config" toml:"isBackupConfig" yaml:"isBackupConfig"`          // 是否从备份区获取配置。默认为 false
	BackupConfigPath  string   `json:"backup_config_path" toml:"backupConfigPath" yaml:"backupConfigPath"`    // 备份配置文件路径
}

type ParseMethodType string

// 解析配置
const (
	ParseMethodTypeViper ParseMethodType = "viper"
)
