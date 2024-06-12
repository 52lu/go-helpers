package example

import (
	"context"
	"fmt"
	"gitlab.dev.olanyun.com/olanyun/saas-utils/cryptoutil/gmcrypt"
	"testing"
)

var (
	// 私钥
	_privateECDSAKey = "JShsBOJL0RgPAoPttEB1hgtPAvCikOl0V1oTOYL7k5U="
	_publicECDSAKey  = "BE8d9gaaCGrE4dHErWCjqyahm6X8l6Rd7fOGx0gNyrGPp0XDoPbbpu1pk9A2fZ9rEsBtwB1Aecnto/gH4h+T7cY="
	ctx              = context.Background()
)

/*
* @Description: 签名生成
* @Author: LiuQHui
* @Param t
* @Date 2023-11-28 11:06:00
 */
func TestSM2SignGenerate(t *testing.T) {
	forSignStr := "1234567890"
	apiSign, err := gmcrypt.SM3WithSM2Sign(ctx, _privateECDSAKey, forSignStr, nil)
	fmt.Println("原字符串:", forSignStr)
	fmt.Println("生成签名:", apiSign)
	fmt.Println("err:", err)
}

/*
* @Description: 签名验证
* @Author: LiuQHui
* @Param t
* @Date 2023-11-28 11:33:37
 */
func TestSM2SignVerify(t *testing.T) {
	msg := "1234567890"
	sign := "1c3010538e7a5f114ddcdbc1f6036c87c033ae40c516eb155019bb3ecaf7e58eabc7493fdfab4785d14b7ccdfc3ef2a54d9f2921b242ca7920c7aabe965c6739"
	// 验签
	sm2Verify, err := gmcrypt.SM3WithSM2Verify(ctx, _publicECDSAKey, sign, msg, nil)
	fmt.Println("verify:", sm2Verify)
	fmt.Println("err:", err)
}
