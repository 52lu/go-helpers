package bjcatypes

type CaConfig struct {
	Url            string
	AppName        string
	ImplType       string // 实现方式
	BJCACertSn     string // bjca证书序列号
	BJCABase64Cert string // bjca证书
	RedisConfig    *RedisConfig
	IsTest         bool
}

type RedisConfig struct {
	Host string
	Pass string
}
