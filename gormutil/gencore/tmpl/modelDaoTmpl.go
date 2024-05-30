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
	query       query.I{{.ModelName}}Do
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
		query:   query.NewDaoQueryWithHookSession(ctx).{{.ModelName}},
	}{{else}}    return {{.DaoName}}{
		ctx:     ctx,
		connect: query.NewDefaultConnect(ctx),
		query:   query.NewDaoQuerySession(ctx).{{.ModelName}},
	}{{end}}
}

/*
* @Description: 事务{{.DaoName}}
* @Author: gorm.io/gen
* @Param tx
* @Return {{.DaoName}}
* @Date {{.DateTime}}
 */
func ({{.ReceiverPre}} {{.DaoName}}) GetTransQueryDao(tx *query.Query) {{.DaoName}}  {
	return {{.DaoName}}{
		ctx:   {{.ReceiverPre}}.ctx,
		query: tx.WithContext({{.ReceiverPre}}.ctx).{{.ModelName}},
	}
}
`
