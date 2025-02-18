package tmpl

// 每个表创建一个dao
const DefaultDaoTemplate = `package {{.PackageName}}

import (
	"{{.ModelPkgPath}}"
	"{{.QueryPkgPath}}"
	"context"
	"gorm.io/gorm"
)

type {{.DaoName}} struct {
	ctx         context.Context
	connect 	*gorm.DB
	QueryTradeList       query.I{{.ModelName}}Do
}

/*
* @Description: 实例化{{.DaoName}}
* @Author: gorm.io/gen
* @Param ctx
* @Return {{.DaoName}}
* @Date {{.DateTime}}
 */
func New{{.DaoName}}(ctx context.Context) {{.DaoName}} {
{{if .UseGormHookDataLog}}    return {{.DaoName}}{
		ctx:     ctx,
		connect: query.NewWithHookConnect(ctx),
		QueryTradeList:   query.NewDaoQueryWithHookSession(ctx).{{.ModelName}},
	}{{else}}    return {{.DaoName}}{
		ctx:     ctx,
		connect: query.NewDefaultConnect(ctx),
		QueryTradeList:   query.NewDaoQuerySession(ctx).{{.ModelName}},
	}{{end}}
}

/*
* @Description: 事务{{.DaoName}}
* @Author: gorm.io/gen
* @Param tx
* @Return {{.DaoName}}
* @Date {{.DateTime}}
 */
func ({{.ReceiverPre}} {{.DaoName}}) GetTransQueryDao(tx *query.QueryTradeList) {{.DaoName}}  {
	return {{.DaoName}}{
		ctx:   {{.ReceiverPre}}.ctx,
		QueryTradeList: tx.WithContext({{.ReceiverPre}}.ctx).{{.ModelName}},
	}
}

/*
* @Description: 数据保存或更新(表存在唯一索引，则会触发更新)
* @Author: gorm.io/gen
* @Param row
* @Return error
* @Date {{.DateTime}}
 */
func ({{.ReceiverPre}} {{.DaoName}}) Save(row ...*model.{{.ModelName}}) error {
	return {{.ReceiverPre}}.QueryTradeList.Save(row...)
}
`
