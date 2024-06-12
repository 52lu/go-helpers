package bjcatypes

type CommonRequest struct {
	AppName          string `json:"appName,omitempty" xml:"appName"`
	Base64EncodeCert string `json:"base64EncodeCert,omitempty" xml:"base64EncodeCert"`
	BJCACertSn       string `json:"certSn,omitempty" xml:"certSn"`
}

type EncodeEnvelopedDataReq struct {
	CommonRequest
	Cert   string `json:"cert" xml:"cert"`       // 待加密数据
	InData string `json:"oriData" xml:"oriData"` // 待加密数据
}

type DecodeEnvelopedDataReq struct {
	CommonRequest
	InData string `json:"encData" xml:"encData"` // 待解密数据
}

type SignDataByP7AttachReq struct {
	CommonRequest
	InData string `json:"oriData" xml:"oriData"`
}

type VerifySignedDataByP7AttachReq struct {
	CommonRequest
	Pkcs7SignData string `json:"pkcs7SignData" xml:"pkcs7SignData"`
}
