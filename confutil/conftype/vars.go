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
	ApolloConf  *ApolloConfig   // apollo配置
}

type ApolloConfig struct {
	ServiceUrl string   `json:"service_url"` // apollo服务地址
	Cluster    string   `json:"cluster"`     // 集群
	AppId      string   `json:"app_id"`      // appId
	Namespaces []string `json:"namespaces"`  // 命名空间
}

type ParseMethodType string

// 解析配置
const (
	ParseMethodTypeViper ParseMethodType = "viper"
)
