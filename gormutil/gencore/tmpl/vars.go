package tmpl

type BaseDaoTemplateVar struct {
	PackageName string // 包名称
	DateTime    string // 时间
}

type ModelDaoVar struct {
	UseGormHookDataLog bool   // 开启钩子
	PackageName        string // 包名称
	DateTime           string // 时间
	ModelPkgPath       string // model path
	QueryPkgPath       string // query path
	DaoName            string // dao名称
	ModelName          string // 名称
	ReceiverPre        string // 接收器前缀
}
