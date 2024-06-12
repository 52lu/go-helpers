package verifyutil

import (
	"fmt"
	cockroachdbError "github.com/cockroachdb/errors"
	"github.com/go-playground/locales/zh"
	ut "github.com/go-playground/universal-translator"
	"github.com/go-playground/validator/v10"
	zhs "github.com/go-playground/validator/v10/translations/zh"
	"reflect"
	"regexp"
	"strings"
)

var (
	validate = validator.New()                                  // 实例化验证器
	chinese  = zh.New()                                         // 获取中文翻译器
	uni      = ut.New(chinese, chinese)                         // 设置成中文翻译器
	trans, _ = uni.GetTranslator("zh")                          // 获取翻译字典
	_        = zhs.RegisterDefaultTranslations(validate, trans) // 注册翻译器
)

func init() {
	// 注册一个获取json tag的自定义方法
	validate.RegisterTagNameFunc(func(field reflect.StructField) string {
		splitN := strings.SplitN(field.Tag.Get("json"), ",", 2)
		n := splitN[0]
		if n == "-" {
			return ""
		}
		return n
	})
	// 注册额外验证规则
	registerCustomer(validate)
}

/**
*  ValidateStruct
*  @Desc：验证结构体
*  @Author LiuQingHui
*  @param s
*  @return string
*  @Date 2021-11-21 18:05:58
**/
func ValidateStruct(s interface{}) error {
	err := validate.Struct(s)
	if err != nil {
		if errors, ok := err.(validator.ValidationErrors); ok {
			var fieldMap = make(map[string]string, len(errors))
			for _, fieldError := range errors {
				field := fieldError.StructField()
				namespace := fieldError.Namespace()
				fieldMap[namespace] = field
			}
			translate := errors.Translate(trans)
			var errMsg string
			typeOf := reflect.TypeOf(s)
			for key, val := range translate {
				// 字段翻译
				val = transField(typeOf, fieldMap, key, val)
				//val = strings.ReplaceAll(val, "格式必须是", "格式有误~")
				if errMsg == "" {
					errMsg = val
				} else {
					errMsg = fmt.Sprintf("%v、%v", errMsg, val)
				}
			}
			return cockroachdbError.New(errMsg)
		}
	}
	return nil
}

/*
 * transField
 * @Description: 翻译字段信息
 * @Author: LiuQHui
 * @Param typeOf
 * @Param fieldInfo
 * @Param tranErrMsg
 * @Return string
 */
func transField(typeOf reflect.Type, fieldMap map[string]string, fieldInfo, tranErrMsg string) string {
	var tagValue string
	var fieldValue string
	if field, ok := fieldMap[fieldInfo]; ok {
		structAttr := strings.Split(fieldInfo, ".")
		if len(structAttr) == 2 {
			fieldValue = structAttr[1]
		}
		if typeOf.Kind() == reflect.Ptr {
			typeOf = typeOf.Elem()
		}
		// 取出struct对应的tag
		nameAttr, b := typeOf.FieldByName(field)
		if b {
			tagValue = nameAttr.Tag.Get("remark")
		}
	}
	// 拼凑错误信息
	if tagValue != "" && fieldValue != "" {
		tranErrMsg = strings.ReplaceAll(tranErrMsg, fieldValue, tagValue)
	}
	return tranErrMsg
}

/**
 *  @Description: 注册自定义验证规则
 *  @Author: LiuQHui
 *  @param vl
 **/
func registerCustomer(vl *validator.Validate) {
	// 注册校验手机号
	registerPhoneValidate(vl)
}

/**
 *  @Description: 校验手机号
 *  @Author: LiuQHui
 *  @param vl
 **/
func registerPhoneValidate(vl *validator.Validate) {
	// 注册标签
	_ = vl.RegisterValidation("phone", func(fl validator.FieldLevel) bool {
		value := fl.Field().String()
		regular := "^((13[0-9])|(14[5,7])|(15[0-3,5-9])|(17[0,3,5-8])|(18[0-9])|166|198|199|(147))\\d{8}$"
		reg := regexp.MustCompile(regular)
		return reg.MatchString(value)
	})
	// 注册翻译器
	_ = vl.RegisterTranslation("phone", trans, func(ut ut.Translator) error {
		return ut.Add("phone", "{0}格式不正确", true)
	}, func(ut ut.Translator, fe validator.FieldError) string {
		t, _ := ut.T("phone", fe.Field())
		return t
	})
}
