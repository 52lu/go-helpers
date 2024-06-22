package hooks

import (
	"context"
	"fmt"
	"github.com/52lu/go-helpers/gormutil/gormhook/hooktype"
	"github.com/52lu/go-helpers/jsonutil"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
	"log"
	"reflect"
	"strconv"
	"strings"
)

type commonHook struct {
	HookPluginConf hooktype.HookPluginConf
}

var (
	jsonUtil   = jsonutil.Json
	_operateId int64
)

// 变更类型：1增 2改 3删
const (
	RowChangeTypeCreate = 1
	RowChangeTypeUpdate = 2
	RowChangeTypeDelete = 3
)

/*
* @Description: 获取数据id
* @Author: LiuQHui
* @Receiver h
* @Param tx
* @Date 2024-04-09 15:24:22
 */
func (h commonHook) getDataId(tx *gorm.DB) int64 {
	if tx.Statement.ReflectValue.Kind() == reflect.Slice {
		sliceLen := tx.Statement.ReflectValue.Len()
		if sliceLen == 1 {
			// 获取切片中第一个元素
			firstElement := tx.Statement.ReflectValue.Index(0)
			// 如果第一个元素是指针类型，获取其指向的实际值
			if firstElement.Kind() == reflect.Ptr {
				firstElement = firstElement.Elem()
			}
			// 获取第一个元素的ID字段值
			idField := firstElement.FieldByName("ID")
			if idField.IsValid() {
				return idField.Int()
			}
		}
		// 如果是切片，返回0，因为我们不知道应该使用哪个ID
		return 0
	}
	v := tx.Statement.ReflectValue.FieldByName("ID")
	switch v.Kind() {
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		return v.Int()
	case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64:
		return int64(v.Uint())
	default:
		return 0
	}
}

/*
* @Description: 设置操作人id
* @Author: LiuQHui
* @Param ctx
* @Param operateId
* @Date 2024-06-22 20:03:30
 */
func SetOperateId(ctx context.Context, operateId int64) {
	_operateId = operateId
}

/*
* @Description: 获取操作人id信息
* @Author: LiuQHui
* @Receiver h
* @Param ctx
* @Date 2024-04-09 15:26:17
 */
func getOperateId(ctx context.Context) int64 {
	return _operateId
}

/*
* @Description: 获取查询条件
* @Author: LiuQHui
* @Receiver r
* @Param db
* @Date 2024-04-03 17:38:38
 */
func (h commonHook) getConditions(db *gorm.DB) []clause.Expression {
	var conditions []clause.Expression
	clauses := db.Statement.Clauses
	for _, tmp := range clauses {
		w, ok := tmp.Expression.(clause.Where)
		if !ok {
			continue
		}
		conditions = append(conditions, w.Exprs...)
	}
	return conditions
}

/*
* @Description: 获取变更前数据
* @Author: LiuQHui
* @Param tx
* @Date 2024-04-03 12:27:03
 */
func (h commonHook) getChangeBeforeData(tx, dbConnect *gorm.DB) (int64, []map[string]interface{}) {
	//ctx := tx.Statement.Context
	var id int64
	var beforeDataList []map[string]interface{}
	// 获取查询条件
	conditions := h.getConditions(tx)
	if len(conditions) == 0 {
		return id, beforeDataList
	}
	// 拼成SQL
	stmt := &gorm.Statement{
		DB:      tx,
		Table:   tx.Statement.Table,
		Schema:  tx.Statement.Schema,
		Clauses: map[string]clause.Clause{},
	}
	for _, expression := range conditions {
		expression.Build(stmt)
		_, _ = stmt.WriteString(" AND ")
	}
	whereSQL := stmt.SQL.String()
	whereSQL = strings.TrimRight(whereSQL, "AND ")
	execSQL := fmt.Sprintf("SELECT * FROM %s WHERE %s", tx.Statement.Table, whereSQL)
	// 查询数据
	rows, err := dbConnect.Session(&gorm.Session{}).Raw(execSQL, stmt.Vars...).Rows()
	defer rows.Close()
	if err != nil {
		return id, beforeDataList
	}

	for rows.Next() {
		// 创建一个新的 map 用于存储当前行的数据
		rowData := make(map[string]interface{})
		// 将当前行的数据扫描到 map 中
		err = dbConnect.ScanRows(rows, &rowData)
		if err != nil {
			//logger.Warnf(ctx, "Failed to scan row: %v", err)
			continue
		}
		// 将当前行的 map 添加到结果集中
		beforeDataList = append(beforeDataList, rowData)
	}
	// 为空处理
	if len(beforeDataList) == 0 {
		return id, beforeDataList
	}
	if len(beforeDataList) == 1 {
		if v, e := beforeDataList[0]["id"]; e {
			id, _ = strconv.ParseInt(fmt.Sprintf("%v", v), 10, 64)
		}
	}
	return id, beforeDataList
}

/*
* @Description: 钩子通用逻辑处理
* @Author: LiuQHui
* @Param fn
* @Param tx
* @Date 2024-04-09 15:10:26
 */
func (h commonHook) execHookFunc(fn func(tx *gorm.DB), tx *gorm.DB) {
	//ctx := tx.Statement.Context
	// 异常捕获
	defer func() {
		if r := recover(); r != nil {
			log.Printf("数据变更记录异常：%v", r)
		}
	}()
	// 哪些表变更不处理
	if tx.Statement.Table == hooktype.TableNameDataChangeLogModel {
		return
	}
	if strings.Contains(tx.Statement.Table, "_log") {
		return
	}
	// 执行具体逻辑
	fn(tx)
}
