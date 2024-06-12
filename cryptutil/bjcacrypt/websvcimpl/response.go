package websvcimpl

import "encoding/xml"

// 证书获取返回
type GetCertResponse struct {
	XMLName xml.Name `xml:"Envelope"`
	Body    struct {
		GetCertResponse GetCertDataNode `xml:"GetCertResponse"`
	} `xml:"Body"`
}
type GetCertDataNode struct {
	GetCert string `xml:"GetCert"`
}

// 加密返回
type EncodeEnvelopedDataResponse struct {
	XMLName xml.Name `xml:"Envelope"`
	Body    struct {
		EncodeEnvelopedDataResponse EncodeEnvelopedDataNode `xml:"EncodeEnvelopedDataResponse"`
	} `xml:"Body"`
}
type EncodeEnvelopedDataNode struct {
	EncodeEnvelopedData string `xml:"EncodeEnvelopedData"`
}

// 解密密返回
type DecodeEnvelopedDataResponse struct {
	XMLName xml.Name `xml:"Envelope"`
	Body    struct {
		DecodeEnvelopedDataResp DecodeEnvelopedDataNode `xml:"DecodeEnvelopedDataResponse"`
	} `xml:"Body"`
}
type DecodeEnvelopedDataNode struct {
	DecodeEnvelopedData string `xml:"DecodeEnvelopedData"`
}

// 签名返回
type SignDataByP7AttachResponse struct {
	XMLName xml.Name `xml:"Envelope"`
	Body    struct {
		SignDataByP7AttachResp SignDataByP7AttachNode `xml:"SignDataByP7AttachResponse"`
	} `xml:"Body"`
}
type SignDataByP7AttachNode struct {
	SignDataByP7AttachResult string `xml:"SignDataByP7AttachResult"`
}

// 验证签名返回
type VerifySignedDataByP7AttachResponse struct {
	XMLName xml.Name `xml:"Envelope"`
	Body    struct {
		VerifySignedDataByP7AttachResp VerifySignedDataByP7AttachNode `xml:"VerifySignedDataByP7AttachResponse"`
	} `xml:"Body"`
}
type VerifySignedDataByP7AttachNode struct {
	VerifySignedDataByP7Attach bool `xml:"VerifySignedDataByP7Attach"`
}
