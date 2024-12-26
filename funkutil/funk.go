package funkutil

import "github.com/thoas/go-funk"

/*
 * @Description: 从切片集中获取指定key的切片
 * @Author: LiuQHui
 * @Param data
 * @Param key
 * @Return []int64
**/
func FunkGetInt64(data interface{}, key string) []int64 {
	var result []int64
	tmp := funk.Get(data, key)
	if tmp != nil {
		result = append(result, tmp.([]int64)...)
	}
	return result
}

/*
* @Description: 查询指定字段,返回字符串切片
* @Author: LiuQHui
* @Param data
* @Param key
* @Return []string
* @Date 2024-12-25 18:06:48
 */
func FunkGetString(data interface{}, key string) []string {
	var result []string
	tmp := funk.Get(data, key)
	if tmp != nil {
		result = append(result, tmp.([]string)...)
	}
	return result
}
