package verifyutil

import (
	"regexp"
)

/*
* @Description: 验证手机号合法
* @Author: LiuQHui
* @Param phone
* @Return bool
* @Date 2023-07-21 12:06:45
 */
func VerifyPhone(phone string) bool {
	// 匹配规则
	// ^1第一位为一
	// [345789]{1} 后接一位345789 的数字
	// \\d \d的转义 表示数字 {9} 接9位
	// $ 结束符
	regRuler := "^1[3456789]{1}\\d{9}$"
	// 正则调用规则
	reg := regexp.MustCompile(regRuler)
	// 返回 MatchString 是否匹配
	return reg.MatchString(phone)
}
