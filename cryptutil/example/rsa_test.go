/**
 * @Author Mr.LiuQH
 * @Description RSA测试使用
 * @Date 2021/7/1 4:40 下午
 **/
package example

import (
	"fmt"
	"gitlab.dev.olanyun.com/olanyun/saas-utils/cryptoutil"
	"testing"
)

const (
	keyPath           = "./tmp"
	publicPKCS8KeyStr = `
-----BEGIN RSA PUBLIC KEY-----
MIGJAoGBANXpPuaMUTXeR5xwg93AI8A8fMx+lUvAHnPeSHeH8lM71UYvhF2065aL
DYZercxOdeQyOwZMo4oo3dJ7p+5kKoKWmeZjR/TIwYQwoucYdXVdfu1o374ecuVs
/dOl7z57oLk+f31VRTv8hs3nW0I0ymiNqIVfwPVdcs+wAjB0iHPjAgMBAAE=
-----END RSA PUBLIC KEY-----
`
	privatePKCS1KeyStr = `
-----BEGIN RSA PRIVATE KEY-----
MIICXQIBAAKBgQDV6T7mjFE13keccIPdwCPAPHzMfpVLwB5z3kh3h/JTO9VGL4Rd
tOuWiw2GXq3MTnXkMjsGTKOKKN3Se6fuZCqClpnmY0f0yMGEMKLnGHV1XX7taN++
HnLlbP3Tpe8+e6C5Pn99VUU7/IbN51tCNMpojaiFX8D1XXLPsAIwdIhz4wIDAQAB
AoGAPowsK0eeO20abWDs/fs/uOc26qicyQCNJv45HFJxBet56kLcpnXPaT6Rntey
ZAoRKL+TSL1CGJTof7JOoUNCtlFRDsYEO5ABCdZI7IHNUAXL26Jx8cnVd17YMfDa
gGxqSFpZG4b6eWgLGx+h8MkbQMpfc3N2cFVgJGYRng88dRkCQQD7yLH6NgUqafZ7
y6gGcLVKNZX4giPIBoKJubpBsa6YXxlPj3Rn5/qM7VFhkqE7VunAlIbRSa7YMjrv
0YvQYfmdAkEA2X40HZPs6rUz6m1g4siEPRNFiKGfyoKnM7GGuLsRH+I+eG+8Z78s
KU+D1WBQx3zZGZ9U9tQY2SrqaAUXQC1rfwJBANi4JMlzmfqp/mkMIPJ6LPFVMmMW
0WmogM+/N5y4LcolgQnENrQBLt4Cn1vW9ES5SLZkoa6fN4oLokMuIKQa0NkCQE+S
9ypjNulgxs/cmPggeRGHfYdR6w7C4r3tE+d+ufM6abTS3NHwhg3PQ+LLzIJQUXYo
b4OncjfylbTdN/aJJ60CQQCMomuBmw+733Dc+K/wp/+hz/YrEVAP7P7hEK2yjec+
BQrvG4ZRH5rCnPXMa8FksFelAit5UlAX7uybirbkAvV5
-----END RSA PRIVATE KEY-----
`
)

// 测试生成密钥对
func TestGenerateKey(t *testing.T) {
	key, err := cryptoutil.GenerateRSAPKCS1Key(1024, keyPath)
	if err != nil {
		t.Error(err)
	}
	fmt.Printf("%+v\n", key)
}

// 读取密钥
func TestReadKey(t *testing.T) {
	// pkcs1格式-私钥
	privatePKCS1KeyPath := keyPath + "/private.pem"
	privatePKCS1Key, err := cryptoutil.ReadRSAPKCS1PrivateKey(privatePKCS1KeyPath)
	if err != nil {
		t.Error(err)
	}
	fmt.Printf("PKCS1私钥: %#v\n", privatePKCS1Key)

	// pkcs8格式-公钥
	publicPKCS8KeyPath := keyPath + "/public.pem"
	publicPKCS8Key, err := cryptoutil.ReadRSAPublicKey(publicPKCS8KeyPath)
	if err != nil {
		t.Error(err)
	}
	fmt.Printf("PKCS8公钥: %#v\n", publicPKCS8Key)
}

func TestReadKeyByStr(t *testing.T) {
	// 私钥
	// pkcs1格式-私钥
	privatePKCS1Key, err := cryptoutil.ReadRSAPKCS1PrivateKeyByStr(privatePKCS1KeyStr)
	if err != nil {
		t.Error(err)
	}
	fmt.Printf("PKCS1私钥: %#v\n", privatePKCS1Key)

	// pkcs8格式-公钥
	publicPKCS8Key, err := cryptoutil.ReadRSAPublicKey(publicPKCS8KeyStr)
	if err != nil {
		t.Error(err)
	}
	fmt.Printf("PKCS8公钥: %#v\n", publicPKCS8Key)
}

// 加密测试
func TestRsaEncrypt(t *testing.T) {
	publicKeyPath := keyPath + "/tmp/public_ssl.pem"
	data := "123456"
	encrypt, err := cryptoutil.RSAEncrypt(data, publicKeyPath)
	if err != nil {
		t.Error(err)
	}
	fmt.Printf("加密结果:%v \n", encrypt)
}

// 解密测试
func TestRsaDecrypt(t *testing.T) {
	privateKeyPath := keyPath + "/tmp/private_ssl.pem"
	data := "pUYa4set6XkBshfio5g2hzPx1tA67sxEvJBpJiuK3McJ9cPJAXzuRkWIy4s6cDQOhrPUaNXhr3M3WLHH19/eaqcNZz1yOFZwgGKmkWtdmygtLB/wrDant9uRfXrvzlV9iMq+cUlqsrwuCa0wcGEBNHRhIJOQSTs+SxaRTeoRCbU="
	encrypt, err := cryptoutil.RSADecrypt(data, privateKeyPath)
	if err != nil {
		t.Error(err)
	}
	fmt.Printf("解密结果:%v \n", encrypt)
}

// 加密测试
func TestRsaEncryptByStrKey(t *testing.T) {
	data := "123456"
	encrypt, err := cryptoutil.RSAEncryptByStrKey(data, publicPKCS8KeyStr)
	if err != nil {
		t.Error(err)
	}
	fmt.Printf("加密结果:%v \n", encrypt)
}

// 解密测试
func TestRsaDecryptByStrKey(t *testing.T) {
	data := "ffJ5/1dCFHvSTkE66MWKfia+FNSZQhriJ/N+LwgcSJYXHKRvotSm2Fgf3YTXNaBpaKQgqOue837wAvLTOjwbSqGAP3CQAIiky4J4Sgny2xa1AJ0iGQ0o7Rj1j2/qzN4ywUlSDVphY5u3/bP8wc6Xoc+JiGXPEOZEvPx9VjS88OQ="
	encrypt, err := cryptoutil.RSADecryptByStrKey(data, privatePKCS1KeyStr)
	if err != nil {
		t.Error(err)
	}
	fmt.Printf("解密结果:%v \n", encrypt)
}

// 数据加签
func TestAddSign(t *testing.T) {
	privateKeyPath := keyPath + "/tmp/private_ssl.pem"
	data := "123456"
	sign, err := cryptoutil.GetRSASign(data, privateKeyPath)
	if err != nil {
		t.Error(err)
	}
	fmt.Printf("数据签名: %v \n", sign)
}

// 数据签名验证
func TestVaSign(t *testing.T) {
	publicKeyPath := keyPath + "/tmp/public_ssl.pem"
	data := "123456"
	sign := "QnGqGbIqoHjJG1l+JiaOKWBdX+h00lnKCoO2rTYKIro9hoaDj7nqmu+Mxsuo+2jumicvCNBZNOpMzYryjZf0x7Q4ycLBtqtCWuFRasiInUO7Avy19LRTjdMf2xw9968vilB/xEAQ53JXIDUVvCsMxTfpHI9oRiWEGXWNkhfkjkQ="
	verifyRsaSign, err := cryptoutil.VerifyRsaSign(data, publicKeyPath, sign)
	if err != nil {
		fmt.Printf("验签失败: %v \n", err)
	}
	fmt.Printf("验签结果: %v \n", verifyRsaSign)
}
