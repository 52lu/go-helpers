package tmpl

// dao通用方法(表有id的方法)
const DaoCommonByIdMethod = `

/*
* @Description: 根据id更新字段，value可以是map[string]interface{}|model
* @Author: gorm.io/gen
* @Receiver {{.ReceiverPre}}
* @Param id
* @Param values
* @Return error
* @Date {{.DateTime}}
 */
func ({{.ReceiverPre}} {{.DaoName}}) UpdateById(id int64,values interface{}) error  {
	_, err := {{.ReceiverPre}}.query.Where(query.{{.ModelName}}.ID.Eq(id)).Updates(values)
	return err
}

/*
* @Description: 根据id查询
* @Author: gorm.io/gen
* @Receiver {{.ReceiverPre}}
* @Param id
* @Return *model.{{.ModelName}}
* @Return error
* @Date {{.DateTime}}
 */
func ({{.ReceiverPre}} {{.DaoName}}) FindById(id int64) (*model.{{.ModelName}}, error)  {
	return {{.ReceiverPre}}.query.FindById(id)
}
`
