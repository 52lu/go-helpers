package verifyutil

import (
	"fmt"
	"testing"
)

/*
* @Description: 解析身份证信息
* @Author: LiuQHui
* @Param t
* @Date 2024-06-12 11:19:23
 */
func TestParseIdCard(t *testing.T) {
	idCard := "341221198903048135"
	card, err := ParseIdCard(idCard)
	fmt.Println(card)
	fmt.Println(err)
}

type UserParam struct {
	Name string `json:"name" validate:"required" remark:"姓名"`
	Age  int    `json:"age" validate:"required,gte=18" remark:"年龄"`
}

/*
* @Description: 校验结构体
* @Author: LiuQHui
* @Param t
* @Date 2024-06-12 11:23:50
 */
func TestValidateStruct(t *testing.T) {
	tmp := UserParam{
		Name: "",
		Age:  0,
	}
	err := ValidateStruct(tmp)
	if err != nil {
		t.Error(err)
		return
	}
	fmt.Println("successfully")
}
