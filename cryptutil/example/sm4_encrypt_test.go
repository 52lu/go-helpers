package example

import (
	"fmt"
	"gitlab.dev.olanyun.com/olanyun/saas-utils/cryptoutil/gmcrypt"
	"testing"
	"time"
)

// sm4加解密key
var _sm4EncryptKey = "0123456789abcdef"

/*
* @Description: Sm4加密
* @Author: LiuQHui
* @Param t
* @Date 2023-11-28 12:28:08
 */
func TestSmEncrypt(t *testing.T) {
	data := fmt.Sprintf("[1234567890abcdef-abc-333-您好-大家好大家好] -> %v", time.Now().Format(time.DateTime))
	fmt.Println("原数据:", data)
	encrypt, err := gmcrypt.Sm4Encrypt(ctx, _sm4EncryptKey, data)
	fmt.Println("加密结果: ", encrypt)
	fmt.Println("err:", err)
}

/*
* @Description: Sm4解密
* @Author: LiuQHui
* @Param t
* @Date 2023-11-28 12:27:59
 */
func TestSm4Decrypt(t *testing.T) {
	// 解密
	data := "af9c9e4330bcbcf01e47534c7b1cd498ac990553daa6d0113fb82bee2b1304784a59f5f238850dcf9e65e0cc61b43aaf0179e3066f3796986366a1ba2b78cefe14b4b24a979e1b3bc6022f87eb4db8d1"
	fmt.Println("原数据:", data)
	decrypt, err := gmcrypt.Sm4Decrypt(ctx, _sm4EncryptKey, data)
	if err != nil {
		t.Errorf("sm4 dec error:%s", err)
		return
	}
	fmt.Println("解密结果:", decrypt)
}
