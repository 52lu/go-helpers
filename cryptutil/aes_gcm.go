package cryptutil

import (
	"crypto/aes"
	"crypto/cipher"
	"encoding/base64"
	"fmt"
)

/*
* @Description: 加密
* @Author: LiuQHui
* @Param data
* @Param key
* @Return string
* @Date 2024-06-12 14:01:50
 */
func AesEncryptByGCM(data, key string) string {
	block, err := aes.NewCipher([]byte(key))
	if err != nil {
		panic(fmt.Sprintf("NewCipher error:%s", err))
	}
	gcm, err := cipher.NewGCM(block)
	if err != nil {
		panic(fmt.Sprintf("NewGCM error:%s", err))
	}
	// 生成随机因子(这里固定取密钥指定位数)
	//nonce := make([]byte, gcm.NonceSize())
	//if _,err := io.ReadFull(rand.Reader,nonce); err != nil {
	//	panic(fmt.Sprintf("make rand nonce error:%s", err))
	//}
	nonceStr := key[:gcm.NonceSize()]
	nonce := []byte(nonceStr)
	fmt.Printf("nonceStr = %v \n", nonceStr)
	seal := gcm.Seal(nonce, nonce, []byte(data), nil)
	return base64.StdEncoding.EncodeToString(seal)
}

/*
* @Description: 解密
* @Author: LiuQHui
* @Param data
* @Param key
* @Return string
* @Date 2024-06-12 14:02:02
 */
func AesDecryptByGCM(data, key string) string {
	// 反解base64
	dataByte, err := base64.StdEncoding.DecodeString(data)
	if err != nil {
		panic(fmt.Sprintf("base64 DecodeString error:%s", err))
	}
	block, err := aes.NewCipher([]byte(key))
	if err != nil {
		panic(fmt.Sprintf("NewCipher error:%s", err))
	}
	gcm, err := cipher.NewGCM(block)
	if err != nil {
		panic(fmt.Sprintf("NewGCM error:%s", err))
	}
	nonceSize := gcm.NonceSize()
	if len(dataByte) < nonceSize {
		panic("dataByte to short")
	}
	nonce, ciphertext := dataByte[:nonceSize], dataByte[nonceSize:]
	open, err := gcm.Open(nil, nonce, ciphertext, nil)
	if err != nil {
		panic(fmt.Sprintf("gcm Open error:%s", err))
	}
	return string(open)
}
