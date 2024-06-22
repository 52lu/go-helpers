package gormhook

import (
	"github.com/52lu/go-helpers/gormutil/gormhook/hooks"
	"github.com/52lu/go-helpers/gormutil/gormhook/hooktype"
	"gorm.io/gorm"
	"sync"
)

type (
	HookInterface interface {
		AddHooks()
	}
	gormHookPlugin struct{}
)

func newGormHookPlugin() *gormHookPlugin {
	return &gormHookPlugin{}
}

/*
* @Description: 注册插件
* @Author: LiuQHui
* @Receiver g
* @Param db
* @Date 2024-04-09 14:54:36
 */
func (g *gormHookPlugin) register(hookList []HookInterface) {
	// 防止多次注册
	for _, hook := range hookList {
		hook.AddHooks()
	}
}

var (
	onceHookInstance sync.Once
)

/*
* @Description: 获取带有全局钩子的实例
* @Author: LiuQHui
* @Date 2024-04-09 17:10:21
 */
func SetGlobalHookInstance(db *gorm.DB) {
	err := db.Set("gorm:table_options", "ENGINE=InnoDB CHARSET=utf8mb4").
		AutoMigrate(&hooktype.DataChangeLogModel{})
	if err != nil {
		panic(err)
	}
	onceHookInstance.Do(func() {
		hookPluginConf := hooktype.HookPluginConf{
			DB:                   db,
			FilterDiffColumnList: []string{"created_at", "updated_at"},
		}
		hookPlugin := newGormHookPlugin()
		// 设置钩子
		hookList := []HookInterface{
			hooks.NewCreateHook(hookPluginConf),
			hooks.NewUpdateHook(hookPluginConf),
			hooks.NewDeleteHook(hookPluginConf),
		}
		hookPlugin.register(hookList)
	})
}
