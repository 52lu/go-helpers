package gmcrypt

import (
	"context"
	"encoding/hex"
	"fmt"
	"github.com/tjfoc/gmsm/sm4"
)

/*
* @Description: 加密
* @Author: LiuQHui
* @Param ctx
* @Param encryptKey
* @Param data
* @Date 2023-11-28 13:45:02
 */
func Sm4Encrypt(ctx context.Context, encryptKey, data string) (string, error) {
	ecbMsg, err := sm4.Sm4Ecb([]byte(encryptKey), []byte(data), true)
	if err != nil {
		return "", err
	}
	//_encryptData := fmt.Sprintf("%x", ecbMsg)
	return fmt.Sprintf("%x", ecbMsg), nil
}

/*
* @Description: 解密
* @Author: LiuQHui
* @Param ctx
* @Param encryptKey
* @Param data
* @Date 2023-11-28 13:46:20
 */
func Sm4Decrypt(ctx context.Context, encryptKey, data string) (string, error) {
	// 解密
	decodeString, err := hex.DecodeString(data)
	if err != nil {
		return "", err
	}
	//sm4Ecb模式pksc7填充解密
	ecbDec, err := sm4.Sm4Ecb([]byte(encryptKey), decodeString, false)
	if err != nil {
		return "", err
	}
	return string(ecbDec), nil
}
