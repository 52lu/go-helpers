package httpimpl

const (
	ResponseSuccessMsg  = "SUCCESS"
	ResponseSuccessCode = 0
	// 通过sn获取证书
	ApiGetCertBySn = "/api/cert/getCertBySn"
	// 数字信封加密
	ApiEncodeEnvelopedData = "/api/cipher/encodeEnvelopedData"
	// 数字信封解密
	ApiDecodeEnvelopedData = "/api/cipher/decodeEnvelopedData"
	//  PKCS7 签名(带原文)-生成
	ApiSignDataByP7Attach = "/api/pkcs7/signDataByP7Attach"
	//  PKCS7 签名(带原文)-验证
	ApiVerifySignedDataByP7Attach = "/api/pkcs7/verifySignedDataByP7Attach"
)
