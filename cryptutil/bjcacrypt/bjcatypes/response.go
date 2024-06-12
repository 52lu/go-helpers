package bjcatypes

type CommonResp struct {
	Body    interface{} `json:"body" xml:"body"`       // 具体响应
	MsgInfo string      `json:"msgInfo" xml:"msgInfo"` // 成功为 SUCCESS，失败为 FAIL，HTTP 错误 404/500 为对应的描述信息
	Status  int         `json:"status" xml:"status"`   // 0:成功
}

type GetServerCertificateResp struct {
	Base64Cert string `json:"base64Cert" xml:"base64Cert"`
}

// 加密返回
type EncodeEnvelopedDataResp struct {
	EnvelopData string `json:"envelopData" xml:"envelopData"`
}

// 解密返回
type DecodeEnvelopedDataResp struct {
	PlainData string `json:"plainData" xml:"plainData"`
}

// 生成签名
type SignDataByP7AttachResp struct {
	P7SignAttach string `json:"p7SignAttach" xml:"p7SignAttach"`
}

// 验证签名
type VerifySignedDataByP7AttachResp struct {
	VerifyRes bool `json:"verifyRes" xml:"verifyRes"`
}
