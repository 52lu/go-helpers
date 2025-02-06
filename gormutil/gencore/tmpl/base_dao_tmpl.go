package tmpl

// dao基类
const DefaultBaseDaoTemplate = `//tip:生成之后，后面再次运行不会再次覆盖，如需覆盖请手动删除

package {{.PackageName}}

import (
	"context"
	"github.com/52lu/go-helpers/gormutil/gormhook"
	"github.com/52lu/go-helpers/gormutil/gormtype"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"sync"
)

var (
	syncOnceDefault sync.Once
	syncOnceHook    sync.Once
	defaultDB       *gorm.DB
	hookDB          *gorm.DB
	err             error
	querySession    *queryCtx // dao链接会话

)

/*
* @Description: 连接池(带有钩子)
* @Author: gorm.io/gen
* @Param ctx
* @Return *gorm.DB
* @Date {{.DateTime}}
 */
func NewWithHookConnect(ctx context.Context,  connectConfig *gormtype.OrmConnectConfig) *gorm.DB {
	syncOnceHook.Do(func() {
		// 获取连接池
		hookDB, err = getDBConnectPool(connectConfig)
		if err != nil {
			panic(err)
		}
		//全局钩子
		gormhook.SetGlobalHookInstance(hookDB)
	})
	return hookDB.WithContext(ctx)
}

/*
* @Description: 默认连接方式
* @Author: LiuQHui
* @Param ctx
* @Param dsnOption
* @Date {{.DateTime}}
 */
func NewDefaultConnect(ctx context.Context, connectConfig *gormtype.OrmConnectConfig) *gorm.DB {
	syncOnceDefault.Do(func() {
		// 连接数据库
		defaultDB, err = getDBConnectPool(connectConfig)
		if err != nil {
			panic(err)
		}
	})
	return defaultDB.WithContext(ctx)
}

/*
* @Description: dao查询会话(有钩子)
* @Author: LiuQHui
* @Param ctx
* @Param dsnOption
* @Date {{.DateTime}}
 */
func NewDaoQueryWithHookSession(ctx context.Context,connectConfig *gormtype.OrmConnectConfig) *queryCtx {
	SetDefault(NewWithHookConnect(ctx,connectConfig))
	querySession = Q.WithContext(ctx)
	return querySession
}

/*
* @Description: dao查询会话(无钩子)
* @Author: LiuQHui
* @Param ctx
* @Param dsnOption
* @Date {{.DateTime}}
 */
func NewDaoQuerySession(ctx context.Context,connectConfig *gormtype.OrmConnectConfig) *queryCtx {
	SetDefault(NewDefaultConnect(ctx,connectConfig))
	querySession = Q.WithContext(ctx)
	return querySession
}

/*
* @Description: 获取连接池
* @Author: gorm.io/gen
* @Param ctx
* @Return *gorm.DB
* @Date {{.DateTime}}
 */
func getDBConnectPool(connectConfig *gormtype.OrmConnectConfig) (*gorm.DB, error) {
	var mysqlDsn string
	if len(dsnOption) > 0 {
		mysqlDsn = dsnOption[0]
	} else {
		// TODO 结合项目，从配置中获取
		mysqlDsn = ""
	}
	if mysqlDsn == "" {
		panic("未设置MySQL连接配置")
	}

    


	// 连接数据库
	return gorm.Open(mysql.Open(mysqlDsn))
}
`
