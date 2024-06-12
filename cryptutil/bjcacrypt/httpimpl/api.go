package httpimpl

import (
	"context"
	"fmt"
	"gitlab.dev.olanyun.com/olanyun/saas-utils/cryptoutil/bjcacrypt/bjcatypes"
	"time"
)

/*
* @Description: 获取证书
* @Author: LiuQHui
* @Receiver c
* @Param ctx
* @Date 2023-12-01 14:52:12
 */
func (c *BjCaCryptHttpImplClient) GetCertBySn(ctx context.Context) (*bjcatypes.GetServerCertificateResp, error) {
	// 判断证书缓存是否存在&有效
	cacheKey := fmt.Sprintf("CaCert:%v:%v:%v", c.Config.AppName, c.Config.BJCACertSn, time.Now().Format(time.DateOnly))
	var result bjcatypes.GetServerCertificateResp
	if c.rd != nil {
		cacheRes, _ := c.rd.Get(cacheKey)
		if cacheRes != "" {
			err := jsonhelper.UnmarshalFromString(cacheRes, &result)
			return &result, err
		}
	}
	httpResp, err := c.sendHttp(ctx, ApiGetCertBySn, bjcatypes.CommonRequest{
		AppName:    c.Config.AppName,
		BJCACertSn: c.Config.BJCACertSn,
	})
	if err != nil {
		return nil, err
	}
	if c.rd != nil {
		_ = c.rd.SetexCtx(ctx, cacheKey, httpResp, 86400)
	}
	err = jsonhelper.UnmarshalFromString(httpResp, &result)
	return &result, err
}

/*
* @Description: 数字信封加密
* @Author: LiuQHui
* @Receiver c
* @Param ctx
* @Param req
* @Date 2023-12-01 17:18:04
 */
func (c *BjCaCryptHttpImplClient) EncodeEnvelopedData(ctx context.Context, req bjcatypes.EncodeEnvelopedDataReq) (*bjcatypes.EncodeEnvelopedDataResp, error) {
	req.AppName = c.Config.AppName
	req.Cert = c.Config.BJCABase64Cert
	httpResp, err := c.sendHttp(ctx, ApiEncodeEnvelopedData, req)
	if err != nil {
		return nil, err
	}
	var result bjcatypes.EncodeEnvelopedDataResp
	err = jsonhelper.UnmarshalFromString(httpResp, &result)
	return &result, err
}

/*
* @Description: 解密数据
* @Author: LiuQHui
* @Receiver c
* @Param ctx
* @Param req
* @Date 2023-12-01 17:21:42
 */
func (c *BjCaCryptHttpImplClient) DecodeEnvelopedData(ctx context.Context, req bjcatypes.DecodeEnvelopedDataReq) (*bjcatypes.DecodeEnvelopedDataResp, error) {
	req.AppName = c.Config.AppName
	httpResp, err := c.sendHttp(ctx, ApiDecodeEnvelopedData, req)
	if err != nil {
		return nil, err
	}
	var result bjcatypes.DecodeEnvelopedDataResp
	err = jsonhelper.UnmarshalFromString(httpResp, &result)
	return &result, err
}

/*
* @Description: PKCS7签名(带原文)-生成签名
* @Author: LiuQHui
* @Receiver c
* @Param ctx
* @Param req
* @Date 2023-12-01 17:25:46
 */
func (c *BjCaCryptHttpImplClient) SignDataByP7Attach(ctx context.Context, req bjcatypes.SignDataByP7AttachReq) (*bjcatypes.SignDataByP7AttachResp, error) {
	req.AppName = c.Config.AppName
	httpResp, err := c.sendHttp(ctx, ApiSignDataByP7Attach, req)
	if err != nil {
		return nil, err
	}
	var result bjcatypes.SignDataByP7AttachResp
	err = jsonhelper.UnmarshalFromString(httpResp, &result)
	return &result, err
}

/*
* @Description: PKCS7签名(带原文)-验证签名
* @Author: LiuQHui
* @Receiver c
* @Param ctx
* @Param req
* @Date 2023-12-01 17:25:46
 */
func (c *BjCaCryptHttpImplClient) VerifySignedDataByP7Attach(ctx context.Context, req bjcatypes.VerifySignedDataByP7AttachReq) (*bjcatypes.VerifySignedDataByP7AttachResp, error) {
	req.AppName = c.Config.AppName
	httpResp, err := c.sendHttp(ctx, ApiVerifySignedDataByP7Attach, req)
	if err != nil {
		return nil, err
	}
	var result bjcatypes.VerifySignedDataByP7AttachResp
	err = jsonhelper.UnmarshalFromString(httpResp, &result)
	return &result, err
}
