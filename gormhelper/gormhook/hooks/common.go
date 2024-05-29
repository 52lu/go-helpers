package hooks

import (
	"context"
	"cq-partner-api/app/helper"
	"fmt"
	"gitlab.weimiaocaishang.com/components/go-gin/logger"
	"gitlab.weimiaocaishang.com/components/go-utils/wmutil"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
	"reflect"
	"strconv"
	"strings"
)

type commonHook struct {
}

var (
	jsonUtil = wmutil.Json()
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
	return tx.Statement.ReflectValue.FieldByName("ID").Int()
}

/*
* @Description: 获取操作人id信息
* @Author: LiuQHui
* @Receiver h
* @Param ctx
* @Date 2024-04-09 15:26:17
 */
func (h commonHook) getOperateId(ctx context.Context) int64 {
	// 获取后台登录人
	operateId := helper.GetAdminLoginAuthId(ctx)
	if operateId != 0 {
		return operateId
	}
	// 获取C端用户
	token, err := helper.GetApiUserInfoWithToken(ctx)
	if err != nil {
		return 0
	}
	return token.Info.Id
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
	ctx := tx.Statement.Context
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
			logger.Warnf(ctx, "Failed to scan row: %v", err)
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
