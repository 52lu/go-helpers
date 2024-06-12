package bjcacrypt

import (
	"context"
	"gitlab.dev.olanyun.com/olanyun/saas-utils/cryptoutil/bjcacrypt/bjcatypes"
)

type BJCACryptInterface interface {
	// 获取证书
	GetCertBySn(ctx context.Context) (*bjcatypes.GetServerCertificateResp, error)
	//  数字信封加密
	EncodeEnvelopedData(ctx context.Context, req bjcatypes.EncodeEnvelopedDataReq) (*bjcatypes.EncodeEnvelopedDataResp, error)
	// 数字信封解密
	DecodeEnvelopedData(ctx context.Context, req bjcatypes.DecodeEnvelopedDataReq) (*bjcatypes.DecodeEnvelopedDataResp, error)
	// PKCS7签名(带原文)-生成签名
	SignDataByP7Attach(ctx context.Context, req bjcatypes.SignDataByP7AttachReq) (*bjcatypes.SignDataByP7AttachResp, error)
	// PKCS7签名(带原文)-验证签名
	VerifySignedDataByP7Attach(ctx context.Context, req bjcatypes.VerifySignedDataByP7AttachReq) (*bjcatypes.VerifySignedDataByP7AttachResp, error)
}
