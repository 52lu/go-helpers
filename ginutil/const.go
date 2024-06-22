package ginutil

// 错误码常量
const (
	RespCodeSuccess       = 200 // 响应成功
	RespCodeErrorNotFound = 404 // 路由不存在
	RespCodeError         = 0   // 处理异常
	RespCodeParam         = -1  // 参数解析异常)
)

// 枚举类型
type RespEnum struct {
	Code int    `json:"code"`
	Msg  string `json:"msg"`
}

var (
	RespEnumSuccess       = RespEnum{Code: RespCodeSuccess, Msg: "处理成功"}
	RespEnumErrorParam    = RespEnum{Code: RespCodeParam, Msg: "参数解析异常"}
	RespEnumErrorNotFound = RespEnum{Code: RespCodeErrorNotFound, Msg: "路由不存在或请求方式错误"}
)

// 运行环境
const (
	GinRunEnvLocal = "local" // 本地运行
	GinRunEnvTest  = "test"  // 测试环境
	GinRunEnvStage = "stage" // 预生产环境
	GinRunEnvProd  = "prod"  // 生产环境
)
