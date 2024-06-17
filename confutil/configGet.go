package confutil

import (
	"time"
)

/*
* @Description: 根据key获取配置值(任意类型)
* @Author: LiuQHui
* @Param key
* @Return interface{}
* @Date 2024-06-04 14:18:30
 */
func Get(key string) interface{} { return _client.impl.Get(key) }

/*
* @Description: 根据key获取配置值(字符串类型)
* @Author: LiuQHui
* @Param key
* @Return interface{}
* @Date 2024-06-04 14:18:30
 */
func GetString(key string) string { return _client.impl.GetString(key) }

/*
* @Description: 根据key获取配置值(不然类型)
* @Author: LiuQHui
* @Param key
* @Return interface{}
* @Date 2024-06-04 14:18:30
 */
func GetBool(key string) bool { return _client.impl.GetBool(key) }

/*
* @Description: 根据key获取配置值(int类型)
* @Author: LiuQHui
* @Param key
* @Return interface{}
* @Date 2024-06-04 14:18:30
 */
func GetInt(key string) int { return _client.impl.GetInt(key) }

/*
* @Description: 根据key获取配置值(int64类型)
* @Author: LiuQHui
* @Param key
* @Return interface{}
* @Date 2024-06-04 14:18:30
 */
func GetInt64(key string) int64 { return _client.impl.GetInt64(key) }

/*
* @Description: 根据key获取配置值(float64类型)
* @Author: LiuQHui
* @Param key
* @Return interface{}
* @Date 2024-06-04 14:18:30
 */
func GetFloat64(key string) float64 { return _client.impl.GetFloat64(key) }

/*
* @Description: 根据key获取配置值(time类型)
* @Author: LiuQHui
* @Param key
* @Return interface{}
* @Date 2024-06-04 14:18:30
 */
func GetTime(key string) time.Time { return _client.impl.GetTime(key) }

/*
* @Description: 根据key获取配置值(int切片)
* @Author: LiuQHui
* @Param key
* @Return interface{}
* @Date 2024-06-04 14:18:30
 */
func GetIntSlice(key string) []int { return _client.impl.GetIntSlice(key) }

/*
* @Description: 根据key获取配置值(字符串切片)
* @Author: LiuQHui
* @Param key
* @Return interface{}
* @Date 2024-06-04 14:18:30
 */
func GetStringSlice(key string) []string { return _client.impl.GetStringSlice(key) }

/*
* @Description: 根据key获取配置值(map)
* @Author: LiuQHui
* @Param key
* @Return interface{}
* @Date 2024-06-04 14:18:30
 */
func GetStringMap(key string) map[string]interface{} { return _client.impl.GetStringMap(key) }

/*
* @Description: 根据key获取配置值(字符串map)
* @Author: LiuQHui
* @Param key
* @Return interface{}
* @Date 2024-06-04 14:18:30
 */
func GetStringMapString(key string) map[string]string { return _client.impl.GetStringMapString(key) }

/*
* @Description: 根据key获取配置值(map)
* @Author: LiuQHui
* @Param key
* @Return interface{}
* @Date 2024-06-04 14:18:30
 */
func GetStringMapStringSlice(key string) map[string][]string {
	return _client.impl.GetStringMapStringSlice(key)
}
