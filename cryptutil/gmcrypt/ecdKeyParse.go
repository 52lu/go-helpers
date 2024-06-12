package gmcrypt

import (
	"context"
	"encoding/base64"
	"encoding/hex"
	"fmt"
	"github.com/tjfoc/gmsm/sm2"
	"github.com/tjfoc/gmsm/x509"
)

func ParseEcStrPrivateKey(ctx context.Context, privateKeyBase64 string) (*sm2.PrivateKey, error) {
	data, err := base64.StdEncoding.DecodeString(privateKeyBase64)
	if err != nil {
		fmt.Println("base64 解码失败:", err)
		return nil, err
	}
	// 将解码后的字节转换为 hex 字符串
	return x509.ReadPrivateKeyFromHex(hex.EncodeToString(data))
}

func ParseEcStrPublicKey(ctx context.Context, privateKeyBase64 string) (*sm2.PublicKey, error) {
	data, err := base64.StdEncoding.DecodeString(privateKeyBase64)
	if err != nil {
		fmt.Println("base64 解码失败:", err)
		return nil, err
	}
	// 将解码后的字节转换为 hex 字符串
	return x509.ReadPublicKeyFromHex(hex.EncodeToString(data))
}
