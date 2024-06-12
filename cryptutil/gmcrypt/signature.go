package gmcrypt

import (
	"bytes"
	"context"
	"crypto/rand"
	"encoding/hex"
	"github.com/tjfoc/gmsm/sm2"
	"math/big"
)

var (
	defaultUid = []byte{0x31, 0x32, 0x33, 0x34, 0x35, 0x36, 0x37, 0x38, 0x31, 0x32, 0x33, 0x34, 0x35, 0x36, 0x37, 0x38}
)

/*
* @Description: 包自带签名，java无法识别
* @Author: LiuQHui
* @Param privateKey
* @Param forSignStr
* @Param uid
* @Date 2023-11-28 11:53:46
 */
func SM3WithSM2Sign(ctx context.Context, privateKeyBase64 string, forSignStr string, uid []byte) (string, error) {
	privateKey, err := ParseEcStrPrivateKey(ctx, privateKeyBase64)
	if err != nil {
		return "", err
	}
	if uid == nil {
		uid = defaultUid
	}
	r, s, err := sm2.Sm2Sign(privateKey, []byte(forSignStr), uid, rand.Reader)
	if err != nil {
		return "", err
	}
	rBytes, sBytes := r.Bytes(), s.Bytes()
	if rLen := len(rBytes); rLen < 32 {
		rBytes = append(make([]byte, 32-rLen), rBytes...)
	}
	if sLen := len(sBytes); sLen < 32 {
		sBytes = append(make([]byte, 32-sLen), sBytes...)
	}

	var buffer bytes.Buffer
	buffer.Write(rBytes)
	buffer.Write(sBytes)
	return hex.EncodeToString(buffer.Bytes()), nil
}

/*
* @Description: 签名验证
* @Author: LiuQHui
* @Param publicKey
* @Param signedHex
* @Param forVerifyStr
* @Param uid
* @Date 2023-11-28 12:04:43
 */
func SM3WithSM2Verify(ctx context.Context, publicKeyBase64 string, signedHex string, forVerifyStr string, uid []byte) (bool, error) {
	publicKey, err := ParseEcStrPublicKey(ctx, publicKeyBase64)
	if err != nil {
		return false, err
	}
	if uid == nil {
		uid = defaultUid
	}
	// 将签名的 hex 字符串解码为字节
	signedBytes, err := hex.DecodeString(signedHex)
	if err != nil {
		return false, err
	}
	// 将要验证的消息和用户标识（UID）转换为字节
	msgBytes := []byte(forVerifyStr)
	// 解析签名中的 r 和 s
	rBytes := signedBytes[:32]
	sBytes := signedBytes[32:]
	r := new(big.Int).SetBytes(rBytes)
	s := new(big.Int).SetBytes(sBytes)

	// 使用 sm2.Sm2Verify 函数验证签名
	verified := sm2.Sm2Verify(publicKey, msgBytes, uid, r, s)
	return verified, nil
}
