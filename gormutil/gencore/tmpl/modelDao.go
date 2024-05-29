package tmpl

// 每个表创建一个dao
const DefaultDaoTemplate = `
package {{.PackageName}}

import (
	"context"
	"{{.ModelPath}}"
	"{{.QueryPath}}"
	"gorm.io/gorm"
)


type {{.DaoName}}Dao struct {
	ctx         context.Context
	conncet 	*gorm.DB
	query       query.IPartnerGoodsModelDo
}

func New{{.DaoName}}Dao(ctx context.Context) {{.DaoName}}Dao {
	return PartnerGoodsDao{
		ctx:         ctx,
		conncet:     query.NewDefaultConn(ctx),
		query:       query.NewDefaultQueryCtx(ctx).{{.ModelName}},
	}
}

func ({{.S}} {{.DaoName}}Dao) GetTransQueryDao(tx *query.Query) {{.DaoName}}Dao  {
	return PartnerGoodsDao{
		ctx:   p.ctx,
		query: tx.WithContext(p.ctx).{{.ModelName}},
	}
}
`
