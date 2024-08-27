package cryptutil

import (
	"crypto/md5"
	"crypto/sha1"
	"fmt"
)

/*
* @Description: 生成md5
* @Author: LiuQHui
* @Param str
* @Return string
* @Date 2024-08-27 11:17:22
 */
func GenerateMd5(str string) string {
	sum := md5.Sum([]byte(str))
	return fmt.Sprintf("%x", sum)
}

/*
* @Description: 生成Sha1
* @Author: LiuQHui
* @Param str
* @Return string
* @Date 2024-08-27 11:18:52
 */
func GenerateSha1(str string) string {
	sum := sha1.Sum([]byte(str))
	return fmt.Sprintf("%x", sum)
}
