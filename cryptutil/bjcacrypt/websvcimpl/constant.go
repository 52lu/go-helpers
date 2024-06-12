package websvcimpl

const (
	// 通过sn获取证书
	ApiGetCert = "GetCert"
	// 数字信封加密
	ApiEncodeEnvelopedData = "EncodeEnvelopedData"
	// 数字信封解密
	ApiDecodeEnvelopedData = "DecodeEnvelopedData"
	// PKCS7 签名(带原文)-生成
	ApiSignDataByP7Attach = "SignDataByP7Attach"
	// PKCS7 签名(带原文)-验证
	ApiVerifySignedDataByP7Attach = "VerifySignedDataByP7Attach"
)
