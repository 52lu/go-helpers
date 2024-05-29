package fileutil

import "os"

/*
* @Description: 判断文件是否存在
* @Author: LiuQHui
* @Param path
* @Date 2024-05-29 14:39:09
 */
func ExistFile(path string) bool {
	_, err := os.Stat(path)
	if err != nil {
		if os.IsExist(err) {
			return true
		}
		return false
	}
	return true
}

/*
* @Description: 目录是否存在
* @Author: LiuQHui
* @Param path
* @Date 2024-05-29 14:39:26
 */
func ExistPath(path string) bool {
	return ExistFile(path)
}

/*
* @Description: 创建目录
* @Author: LiuQHui
* @Param path
* @Date 2024-05-29 14:39:35
 */
func CreatePath(path string) error {
	dirExist := ExistPath(path)
	var err error
	if !dirExist {
		// 不存在则直接创建
		err = os.MkdirAll(path, os.ModePerm)
	}
	return err
}
