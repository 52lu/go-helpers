package pathutil

import (
	"os"
)

/**
 * @Description: 创建目录,(存在则跳过)
 * @Author: LiuQHui
 * @Param d
 * @Return error
 * @Date  2022-11-04 14:13:07
**/
func CreateDir(d string) error {
	if exist, _ := PathIsExist(d); !exist {
		return os.MkdirAll(d, os.ModePerm)
	}
	return nil
}

/**
 * @Description: 目录是否存在
 * @Author: LiuQHui
 * @Param pathName
 * @Return bool
 * @Return error
 * @Date  2022-11-04 14:11:37
**/
func PathIsExist(pathName string) (bool, error) {
	// 获取文件或目录信息
	_, err := os.Stat(pathName)
	// 如果文件或目录存在，则err== nil
	if err == nil {
		return true, nil
	}
	// 使用os.IsNotExist()判断返回的错误类型是否为true,true:说明文件或文件夹不存在
	if os.IsNotExist(err) {
		return false, nil
	}
	// 如果错误类型为其它,则不确定是否在存在
	return false, err
}
