package hooks

import (
	"github.com/52lu/go-helpers/gormutil/gormhook/hooktype"
	"gorm.io/gorm"
)

type deleteHookPlugin struct {
	hookPluginConf hooktype.HookPluginConf
	commonHook
}

func NewDeleteHook(hookPluginConf hooktype.HookPluginConf) *deleteHookPlugin {
	return &deleteHookPlugin{hookPluginConf: hookPluginConf}
}

/*
* @Description: 注册钩子
* @Author: LiuQHui
* @Receiver c
* @Date 2024-04-09 14:53:35
 */
func (d *deleteHookPlugin) AddHooks() {
	processor := d.hookPluginConf.DB.Callback().Delete()
	// 数据删除前触发
	processor.Before("gorm:after_delete").Register("rowAfterDelete", d.rowAfterDelete)
}

/*
* @Description: 删除前触发
* @Author: LiuQHui
* @Receiver d
* @Param tx
* @Date 2024-04-09 15:48:37
 */
func (d *deleteHookPlugin) rowAfterDelete(tx *gorm.DB) {
	execHookFunc(func(tx *gorm.DB) {
		//ctx := tx.Statement.Context
		var changeLog *hooktype.DataChangeLogModel
		var err error
		// 新增逻辑
		changeLog, err = d.formatDeleteRowData(tx)
		if err != nil {
			//logger.Warnf(ctx, "数据格式化处理异常:%v", err)
			return
		}
		// 填充具体执行SQL
		sql := tx.Dialector.Explain(tx.Statement.SQL.String(), tx.Statement.Vars...)
		changeLog.ExecSql = &sql
		// 保存数据
		_ = d.hookPluginConf.DB.Create(changeLog)
	}, tx)
}

/*
* @Description: 格式化删除数控
* @Author: LiuQHui
* @Receiver d
* @Param tx
* @Date 2024-04-09 15:48:04
 */
func (d *deleteHookPlugin) formatDeleteRowData(tx *gorm.DB) (*hooktype.DataChangeLogModel, error) {
	ctx := tx.Statement.Context
	// 修改前数据
	dataId, beforeDataList := d.getChangeBeforeData(tx, d.hookPluginConf.DB)
	var beforeData string
	if len(beforeDataList) == 1 {
		beforeData, _ = jsonUtil.MarshalToString(beforeDataList[0])
	} else {
		beforeData, _ = jsonUtil.MarshalToString(beforeDataList)
	}
	changeLog := &hooktype.DataChangeLogModel{
		DataTable:  tx.Statement.Table,
		DataID:     dataId,
		Type:       RowChangeTypeDelete,
		Before:     &beforeData,
		EffectRows: int64(len(beforeDataList)),
		OperateID:  getOperateId(ctx),
		LogID:      "",
	}
	return changeLog, nil
}
