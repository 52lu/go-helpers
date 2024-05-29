package hooks

import (
	"fmt"
	"github.com/thoas/go-funk"
	"gitlab.weimiaocaishang.com/components/go-gin/logger"
	"gitlab.weimiaocaishang.com/components/go-utils/errutil"

	//"gitlab.weimiaocaishang.com/components/go-gin/logger"
	//"gitlab.weimiaocaishang.com/components/go-utils/errutil"
	"gorm.io/gorm"
	"strings"
	"time"
)

/*
* @Description: 钩子通用逻辑处理
* @Author: LiuQHui
* @Param fn
* @Param tx
* @Date 2024-04-09 15:10:26
 */
func execHookFunc(fn func(tx *gorm.DB), tx *gorm.DB) {
	ctx := tx.Statement.Context
	// 异常捕获
	defer func() {
		if r := recover(); r != nil {
			if conf.IsProdEnv() {
				logger.Warnf(ctx, "数据变更记录异常: %+v", errutil.ThrowErrorMsg(fmt.Sprintf("%v", r)))
			} else {
				logger.Errorf(ctx, "数据变更记录异常: %+v", errutil.ThrowErrorMsg(fmt.Sprintf("%v", r)))
			}
		}
	}()
	// 哪些表变更不处理
	if tx.Statement.Table == model.TableNameDataChangeLogModel {
		return
	}
	if strings.Contains(tx.Statement.Table, "_log") {
		return
	}
	// 执行具体逻辑
	fn(tx)
}

func removeMapZeroValues(inputMap map[string]interface{}, oldMap map[string]interface{}) {
	if len(inputMap) == 0 {
		return
	}
	for key, value := range inputMap {
		// 根据值的类型进行判断，并删除零值
		switch v := value.(type) {
		case int, int8, int16, int32, int64, uint, uint8, uint16, uint32, uint64:
			var noDel bool
			if oldVal, ok := oldMap[key]; ok {
				if !funk.IsEmpty(oldVal) && oldVal != value {
					noDel = true
				}
			}
			if funk.IsEmpty(v) && !noDel {
				delete(inputMap, key)
			}
		case float32, float64:
			var noDel bool
			if oldVal, ok := oldMap[key]; ok {
				if !funk.IsEmpty(oldVal) && oldVal != value {
					noDel = true
				}
			}
			if funk.IsEmpty(v) && !noDel {
				delete(inputMap, key)
			}
		case string:
			var noDel bool
			if oldVal, ok := oldMap[key]; ok {
				if !funk.IsEmpty(oldVal) && oldVal != value {
					noDel = true
				}
			}
			if funk.IsEmpty(v) && !noDel {
				delete(inputMap, key)
			}
		case bool:
			if !v {
				delete(inputMap, key)
			}
		case nil:
			delete(inputMap, key)
		case time.Time:
			if v.IsZero() {
				delete(inputMap, key)
			}
		}
	}
}
