package hooktype

import (
	"time"
)

const TableNameDataChangeLogModel = "data_change_log"

// DataChangeLogModel 数据变更日志
type DataChangeLogModel struct {
	ID         int64     `gorm:"column:id;type:bigint(20) unsigned;primaryKey;autoIncrement:true" json:"id"`
	DataTable  string    `gorm:"column:data_table;type:varchar(128);not null;comment:表名" json:"data_table"`                         // 表名
	DataID     int64     `gorm:"column:data_id;type:bigint(20);not null;comment:数据id" json:"data_id"`                               // 数据id
	Type       int64     `gorm:"column:type;type:tinyint(4);not null;comment:变更类型：1增 2改 3删" json:"type"`                            // 变更类型：1增 2改 3删
	Before     *string   `gorm:"column:before;type:text;comment:修改前" json:"before"`                                                 // 修改前
	After      *string   `gorm:"column:after;type:text;comment:修改后" json:"after"`                                                   // 修改后
	Modified   *string   `gorm:"column:modified;type:text;comment:变更记录" json:"modified"`                                            // 变更记录
	ExecSql    *string   `gorm:"column:exec_sql;type:text;comment:具体执行SQL信息" json:"exec_sql"`                                       // 具体执行SQL信息
	EffectRows int64     `gorm:"column:effect_rows;type:bigint(20) unsigned;not null;comment:影响数据行数" json:"effect_rows"`            // 影响数据行数
	OperateID  int64     `gorm:"column:operate_id;type:bigint(20) unsigned;not null;comment:操作人id" json:"operate_id"`               // 操作人id
	CreatedAt  time.Time `gorm:"column:created_at;type:datetime;not null;default:CURRENT_TIMESTAMP;comment:添加时间" json:"created_at"` // 添加时间
	UpdatedAt  time.Time `gorm:"column:updated_at;type:datetime;not null;default:CURRENT_TIMESTAMP;comment:更新时间" json:"updated_at"` // 更新时间
	LogID      string    `gorm:"column:log_id;type:varchar(128);not null;comment:日志ID" json:"log_id"`                               // 日志ID
}

// TableName DataChangeLogModel's table name
func (*DataChangeLogModel) TableName() string {
	return TableNameDataChangeLogModel
}
