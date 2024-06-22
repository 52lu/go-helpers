package hooks

import (
	"github.com/52lu/go-helpers/gormutil/gormhook/hooktype"
	"gorm.io/gorm"
)

type CreateHookPlugin struct {
	commonHook
	hookPluginConf hooktype.HookPluginConf
}

func NewCreateHook(hookPluginConf hooktype.HookPluginConf) *CreateHookPlugin {
	return &CreateHookPlugin{hookPluginConf: hookPluginConf}
}

/*
* @Description: 注册钩子
* @Author: LiuQHui
* @Receiver c
* @Date 2024-04-09 14:53:35
 */
func (c *CreateHookPlugin) AddHooks() {
	processor := c.hookPluginConf.DB.Callback().Create()
	// 数据保存后触发
	processor.After("gorm:after_create").Register("rowAfterCreate", c.rowAfterCreate)
}

/*
* @Description: 数据保存后触发
* @Author: LiuQHui
* @Receiver c
* @Param tx
* @Date 2024-04-09 15:00:36
 */

func (c *CreateHookPlugin) rowAfterCreate(tx *gorm.DB) {
	execHookFunc(func(tx *gorm.DB) {
		//ctx := tx.Statement.Context
		var changeLog *hooktype.DataChangeLogModel
		var err error
		// 新增逻辑
		changeLog, err = c.formatCreateRowData(tx)
		if err != nil {
			//logger.Warnf(ctx, "数据格式化处理异常:%v", err)
			return
		}
		// 填充具体执行SQL
		sql := tx.Dialector.Explain(tx.Statement.SQL.String(), tx.Statement.Vars...)
		changeLog.ExecSql = &sql
		// 保存数据
		_ = c.hookPluginConf.DB.Create(changeLog)
	}, tx)
}

/*
* @Description: 格式化创建数据
* @Author: LiuQHui
* @Receiver c
* @Param tx
* @Date 2024-04-09 15:26:39
 */
func (c *CreateHookPlugin) formatCreateRowData(tx *gorm.DB) (*hooktype.DataChangeLogModel, error) {
	ctx := tx.Statement.Context
	afterData, _ := jsonUtil.MarshalToString(tx.Statement.Dest)
	var dataMapList []map[string]interface{}
	err := jsonUtil.UnmarshalFromString(afterData, &dataMapList)
	if err != nil {
		return nil, err
	}
	if len(dataMapList) == 1 {
		afterData, _ = jsonUtil.MarshalToString(dataMapList[0])
	}
	// 获取主键
	changeLog := &hooktype.DataChangeLogModel{
		DataTable:  tx.Statement.Table,
		DataID:     c.getDataId(tx),
		Type:       RowChangeTypeCreate,
		After:      &afterData,
		OperateID:  getOperateId(ctx),
		LogID:      "",
		EffectRows: int64(len(dataMapList)),
	}
	return changeLog, nil
}
