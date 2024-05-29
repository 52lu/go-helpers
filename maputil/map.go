package maputil

import "time"

/*
* @Description: 过滤map中的类型零值
* @Author: LiuQHui
* @Param data
* @Date 2024-04-07 16:19:38
 */
func RemoveMapZeroValues(inputMap map[string]interface{}) {
	if len(inputMap) == 0 {
		return
	}
	for key, value := range inputMap {
		// 根据值的类型进行判断，并删除零值
		switch v := value.(type) {
		case int, int8, int16, int32, int64:
			if v == 0 {
				delete(inputMap, key)
			}
		case uint, uint8, uint16, uint32, uint64:
			if v == 0 {
				delete(inputMap, key)
			}
		case float32, float64:
			if v == float64(0) || v == float32(0) {
				delete(inputMap, key)
			}
		case string:
			if v == "" {
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
