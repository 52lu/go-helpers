package websvcimpl

import "encoding/xml"

// 获取证书
type GetCertReq struct {
	XMLName    xml.Name `xml:"tem:GetCert"`
	AppName    string   `xml:"tem:appName"`
	BJCACertSn string   `xml:"tem:sn"`
	//BJCACertSn CharData `xml:"tem:sn"` // 待加密数据
}

// 签名
type SignDataByP7AttachReq struct {
	XMLName xml.Name `xml:"tem:SignDataByP7Attach"`
	AppName string   `xml:"tem:appName"`
	InData  CharData `xml:"tem:inData"`
}

// 签名验证
type VerifySignedDataByP7AttachReq struct {
	XMLName       xml.Name `xml:"tem:VerifySignedDataByP7Attach"`
	AppName       string   `xml:"tem:appName"`
	Pkcs7SignData string   `xml:"tem:pkcs7SignData"`
}

// 加密请求
type EncodeEnvelopedDataReq struct {
	XMLName xml.Name `xml:"tem:EncodeEnvelopedData"`
	AppName string   `xml:"tem:appName"`
	//Base64EncodeCert string   `xml:"tem:base64EncodeCert"`
	Base64EncodeCert CharData `xml:"tem:base64EncodeCert"`
	InData           CharData `xml:"tem:inData"` // 待加密数据
}

// CharData type for handling character data in XML without escaping
type CharData struct {
	XMLName xml.Name
	Value   []byte `xml:",innerxml"`
}
type CharDataString struct {
	XMLName xml.Name
	Value   string `xml:",innerxml"`
}

// 解密
type DecodeEnvelopedDataReq struct {
	XMLName xml.Name `xml:"tem:DecodeEnvelopedData"`
	AppName string   `xml:"tem:appName"`
	InData  CharData `xml:"tem:inData"` // 待解密密数据
}
