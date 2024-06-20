package ginutil

// 响应码类型
type RespEnum struct {
	Code int    `json:"code"`
	Msg  string `json:"msg"`
}

var (
	RespEnumSuccess       = RespEnum{Code: 200, Msg: "处理成功"}
	RespEnumErrorParam    = RespEnum{Code: -1, Msg: "参数解析异常"}
	RespEnumErrorNotFound = RespEnum{Code: 404, Msg: "路由不存在"}
)

var (
	RespSuccess       = 200 // 响应成功
	RespErrorNotFound = 404 // 路由不存在
	RespError         = 0   // 处理异常
)
