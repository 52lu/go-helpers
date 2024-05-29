package hooks

import (
	"52lu/go-helpers/gormutil/gormhook/hooktype"
	"52lu/go-helpers/maputil"
	"github.com/thoas/go-funk"
	"gorm.io/gorm"
)

type updateHookPlugin struct {
	hookPluginConf hooktype.HookPluginConf
	commonHook
}

func NewUpdateHook(hookPluginConf hooktype.HookPluginConf) *updateHookPlugin {
	return &updateHookPlugin{
		hookPluginConf: hookPluginConf,
	}
}

func (u *updateHookPlugin) AddHooks() {
	processor := u.hookPluginConf.DB.Callback().Update()
	// 数据修改后
	processor.Before("gorm:after_update").Register("rowBeforeUpdate", u.rowBeforeUpdate)
}

/*
* @Description: 更新后操作
* @Author: LiuQHui
* @Receiver u
* @Param tx
* @Date 2024-04-09 16:46:36
 */
func (u *updateHookPlugin) rowBeforeUpdate(tx *gorm.DB) {
	execHookFunc(func(tx *gorm.DB) {
		ctx := tx.Statement.Context
		afterData, _ := jsonUtil.MarshalToString(tx.Statement.Dest)
		// 要的更新数据
		var afterDataMap map[string]interface{}
		_ = jsonUtil.UnmarshalFromString(afterData, &afterDataMap)
		// 修改前数据
		dataId, beforeDataList := u.getChangeBeforeData(tx, u.hookPluginConf.DB)
		beforeData, _ := jsonUtil.MarshalToString(beforeDataList)

		var modifiedData string
		// 获取变更数据
		if dataId != 0 && len(beforeDataList) == 1 {
			beforeDataMap := beforeDataList[0]
			// 过滤map中的类型零值
			if len(afterDataMap) == len(beforeDataMap) {
				maputil.RemoveMapZeroValues(afterDataMap)
			} else {
				removeMapZeroValues(afterDataMap, beforeDataMap)
			}
			modifiedData = u.getModifiedData(beforeDataMap, afterDataMap)
		}
		changeLog := &hooktype.DataChangeLogModel{
			DataTable:  tx.Statement.Table,
			DataID:     dataId,
			Type:       RowChangeTypeUpdate,
			Before:     &beforeData,
			EffectRows: int64(len(beforeDataList)),
			After:      &afterData,
			Modified:   &modifiedData,
			OperateID:  u.getOperateId(ctx),
			LogID:      "", //TODO
		}
		// 填充具体执行SQL
		sql := tx.Dialector.Explain(tx.Statement.SQL.String(), tx.Statement.Vars...)
		changeLog.ExecSql = &sql
		// 保存数据
		_ = u.hookPluginConf.DB.Create(changeLog)
	}, tx)
}

/*
* @Description: 获取变更的数据
* @Author: LiuQHui
* @Receiver u
* @Param tmpKey
* @Param afterDataMap
* @Date 2024-04-09 16:28:25
 */
func (u *updateHookPlugin) getModifiedData(beforeDataMap, afterDataMap map[string]interface{}) string {
	var modifiedData string
	// 修改前数据
	if len(beforeDataMap) == 0 || len(afterDataMap) == 0 {
		return modifiedData
	}

	// 变更比较
	var changeDataMap = make(map[string]interface{})
	for k, updateVal := range afterDataMap {
		// 过滤不计较的字段
		if funk.ContainsString(u.hookPluginConf.FilterDiffColumnList, k) {
			continue
		}
		// 进行比较
		if oldVal, ok := beforeDataMap[k]; ok {
			if updateVal != oldVal {
				changeDataMap[k] = updateVal
			}
		}
	}
	modifiedData, _ = jsonUtil.MarshalToString(changeDataMap)
	return modifiedData
}
