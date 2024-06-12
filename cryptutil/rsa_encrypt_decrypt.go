package cryptutil

import (
	"crypto/rand"
	"crypto/rsa"
	"encoding/base64"
)

/**
 * @Description:  加密(使用公钥加密，使用字符串密钥)
 * @Author: LiuQHui
 * @Param data
 * @Param publicKeyStr
 * @Return string
 * @Return error
 * @Date  2023-11-04 14:29:10
**/
func RSAEncryptByStrKey(data, publicKeyStr string) (string, error) {
	// 获取公钥
	rsaPublicKey, err := ReadRSAPublicKeyByStr(publicKeyStr)
	if err != nil {
		return "", err
	}
	// 加密
	encryptPKCS1v15, err := rsa.EncryptPKCS1v15(rand.Reader, rsaPublicKey, []byte(data))
	if err != nil {
		return "", err
	}
	// 把加密结果转成Base64
	encryptString := base64.StdEncoding.EncodeToString(encryptPKCS1v15)
	return encryptString, err
}

/**
 * @Description: 解密(使用私钥解密,使用字符串密钥)
 * @Author: LiuQHui
 * @Param base64data
 * @Param privateKeyPath
 * @Return string
 * @Return error
 * @Date  2023-11-04 14:29:52
**/
func RSADecryptByStrKey(base64data, privateKeyStr string) (string, error) {
	// data反解base64
	decodeString, err := base64.StdEncoding.DecodeString(base64data)
	if err != nil {
		return "", err
	}
	// 读取密钥
	rsaPrivateKey, err := ReadRSAPKCS1PrivateKeyByStr(privateKeyStr)
	if err != nil {
		return "", err
	}
	// 解密
	decryptPKCS1v15, err := rsa.DecryptPKCS1v15(rand.Reader, rsaPrivateKey, decodeString)
	return string(decryptPKCS1v15), err
}

// 加密(使用公钥加密)
func RSAEncrypt(data, publicKeyPath string) (string, error) {
	// 获取公钥
	rsaPublicKey, err := ReadRSAPublicKey(publicKeyPath)
	if err != nil {
		return "", err
	}
	// 加密
	encryptPKCS1v15, err := rsa.EncryptPKCS1v15(rand.Reader, rsaPublicKey, []byte(data))
	if err != nil {
		return "", err
	}
	// 把加密结果转成Base64
	encryptString := base64.StdEncoding.EncodeToString(encryptPKCS1v15)
	return encryptString, err
}

// 解密(使用私钥解密)
func RSADecrypt(base64data, privateKeyPath string) (string, error) {
	// data反解base64
	decodeString, err := base64.StdEncoding.DecodeString(base64data)
	if err != nil {
		return "", err
	}
	// 读取密钥
	rsaPrivateKey, err := ReadRSAPKCS1PrivateKey(privateKeyPath)
	if err != nil {
		return "", err
	}
	// 解密
	decryptPKCS1v15, err := rsa.DecryptPKCS1v15(rand.Reader, rsaPrivateKey, decodeString)
	return string(decryptPKCS1v15), err
}
