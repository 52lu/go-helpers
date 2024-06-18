/**
 * @Author Mr.LiuQH
 * @Description AES加密模式CTR-计算器模式-测试使用
 * @Date 2021/6/29 6:04 下午
 **/
package example

import (
	"fmt"
	"github.com/52lu/go-helpers/cryptutil"
	"strings"
	"testing"
)

// 测试AES-OFB加密
func TestAesEncryptByOFB(t *testing.T) {
	key := strings.Repeat("a", 16)
	data := "123"
	_, base64 := cryptutil.AesEncryptByOFB(data, key)
	_, base642 := cryptutil.AesEncryptByCTR(data, key)
	_, base643 := cryptutil.AesEncryptByCFB(data, key)
	fmt.Printf("加密key: %v \n", key)
	fmt.Printf("加密key长度: %v \n", len(key))
	fmt.Printf("加密数据: %v \n", data)
	fmt.Printf("加密结果(OFB): %v \n", base64)
	fmt.Printf("加密结果(CTR): %v \n", base642)
	fmt.Printf("加密结果(CFB): %v \n", base643)
}

// 测试AES-CTR解密
func TestAesDecryptByOFB(t *testing.T) {
	key := strings.Repeat("a", 16)
	data := "Oe2qKyQC69LyY+8UWLgbVQ=="
	res := cryptutil.AesDecryptByOFB(data, key)
	res1 := cryptutil.AesDecryptByCFB(data, key)
	fmt.Printf("解密key: %v \n", key)
	fmt.Printf("解密数据: %v \n", data)
	fmt.Printf("解密结果(OFB): %v \n", res)
	fmt.Printf("解密结果(CFB): %v \n", res1)
}
