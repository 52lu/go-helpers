package websvcimpl

import (
	"context"
	"encoding/base64"
	"encoding/xml"
	"fmt"
	"github.com/zeromicro/go-zero/core/logx"
	"gitlab.dev.olanyun.com/olanyun/saas-utils/cryptoutil/bjcacrypt/bjcatypes"
	"time"
)

/*
* @Description: 通过sn获取证书
* @Author: LiuQHui
* @Receiver c
* @Param ctx
* @Date 2024-05-08 22:22:22
 */
func (b *BjCaCryptWebserviceImplClient) GetCertBySn(ctx context.Context) (*bjcatypes.GetServerCertificateResp, error) {
	// 判断证书缓存是否存在&有效
	cacheKey := fmt.Sprintf("CaCertBySn:%v:%v:%v", b.Config.AppName, b.Config.BJCACertSn, time.Now().Format("20060102"))
	var result bjcatypes.GetServerCertificateResp
	if b.rd != nil {
		cacheRes, _ := b.rd.Get(cacheKey)
		if cacheRes != "" {
			logx.WithContext(ctx).Infow("从缓存读取证书信息:"+cacheKey, logx.LogField{
				Key:   "cacheRes",
				Value: cacheRes,
			})
			err := jsonhelper.UnmarshalFromString(cacheRes, &result)
			return &result, err
		}
	}
	httpResp, err := b._sendHttp(ctx, ApiGetCert, GetCertReq{
		AppName:    b.Config.AppName,
		BJCACertSn: b.Config.BJCACertSn,
	}, nil)
	if err != nil {
		return nil, err
	}
	// 解析 XML 响应
	var resp GetCertResponse
	if err = xml.Unmarshal([]byte(httpResp), &resp); err != nil {
		logx.WithContext(ctx).Errorf("GetCertBySn Response: %+v", err)
		return nil, err
	}
	// 结果转成base64
	//result.Base64Cert = base64.StdEncoding.EncodeToString([]byte(resp.Body.GetCertResponse.GetCert))
	result.Base64Cert = resp.Body.GetCertResponse.GetCert
	if b.rd != nil {
		certJsonStr, _ := jsonhelper.MarshalToString(result)
		if certJsonStr != "" {
			_ = b.rd.SetexCtx(ctx, cacheKey, certJsonStr, 300)
		}
	}
	return &result, err
}

/*
* @Description: 数字信封加密
* @Author: LiuQHui
* @Receiver b
* @Param ctx
* @Param req
* @Date 2024-05-08 22:24:28
 */
func (b *BjCaCryptWebserviceImplClient) EncodeEnvelopedData(ctx context.Context, req bjcatypes.EncodeEnvelopedDataReq) (*bjcatypes.EncodeEnvelopedDataResp, error) {
	certBytes, err := base64.StdEncoding.DecodeString(b.Config.BJCABase64Cert)
	if err != nil {
		logx.WithContext(ctx).Errorf("证书base64.DecodeString err: %+v", err)
		return nil, err
	}
	param := EncodeEnvelopedDataReq{
		AppName: b.Config.AppName,
		//Base64EncodeCert: b.Config.BJCABase64Cert,
		Base64EncodeCert: CharData{
			Value: certBytes,
		},
		InData: CharData{
			Value: []byte(req.InData),
		},
	}
	//httpResp, err := b._sendHttp(ctx, ApiEncodeEnvelopedData, param, []byte(req.InData))
	httpResp, err := b._sendHttp(ctx, ApiEncodeEnvelopedData, param, map[string][]byte{
		"base64EncodeCert": certBytes,
		"inData":           []byte(req.InData),
	})
	if err != nil {
		return nil, err
	}
	// 解析 XML 响应
	var resp EncodeEnvelopedDataResponse
	if err = xml.Unmarshal([]byte(httpResp), &resp); err != nil {
		logx.WithContext(ctx).Errorf("EncodeEnvelopedData Response: %+v", err)
		return nil, err
	}
	//encodeToString := base64.StdEncoding.EncodeToString([]byte(resp.Body.EncodeEnvelopedDataResponse.EncodeEnvelopedData))
	encodeResult := resp.Body.EncodeEnvelopedDataResponse.EncodeEnvelopedData
	return &bjcatypes.EncodeEnvelopedDataResp{
		EnvelopData: encodeResult,
	}, err
}

/*
* @Description: 数字信封解密
* @Author: LiuQHui
* @Receiver b
* @Param ctx
* @Param req
* @Date 2024-05-08 22:25:09
 */
func (b *BjCaCryptWebserviceImplClient) DecodeEnvelopedData(ctx context.Context, req bjcatypes.DecodeEnvelopedDataReq) (*bjcatypes.DecodeEnvelopedDataResp, error) {
	param := DecodeEnvelopedDataReq{
		AppName: b.Config.AppName,
		//InData:  req.InData,
		InData: CharData{
			Value: []byte(req.InData),
		},
	}
	//httpResp, err := b._sendHttp(ctx, ApiDecodeEnvelopedData, param, []byte(req.InData))
	httpResp, err := b._sendHttp(ctx, ApiDecodeEnvelopedData, param, map[string][]byte{
		"inData": []byte(req.InData),
	})
	if err != nil {
		return nil, err
	}
	// 解析 XML 响应
	var resp DecodeEnvelopedDataResponse
	if err = xml.Unmarshal([]byte(httpResp), &resp); err != nil {
		logx.WithContext(ctx).Errorf("DecodeEnvelopedData Response: %+v", err)
		return nil, err
	}
	return &bjcatypes.DecodeEnvelopedDataResp{
		PlainData: resp.Body.DecodeEnvelopedDataResp.DecodeEnvelopedData,
	}, err
}

/*
* @Description: PKCS7 签名(带原文)-生成
* @Author: LiuQHui
* @Receiver b
* @Param ctx
* @Param req
* @Date 2024-05-08 22:26:50
 */
func (b *BjCaCryptWebserviceImplClient) SignDataByP7Attach(ctx context.Context, req bjcatypes.SignDataByP7AttachReq) (*bjcatypes.SignDataByP7AttachResp, error) {
	param := SignDataByP7AttachReq{
		AppName: b.Config.AppName,
		InData: CharData{
			Value: []byte(req.InData),
		},
	}
	httpResp, err := b._sendHttp(ctx, ApiSignDataByP7Attach, param, map[string][]byte{
		"inData": []byte(req.InData),
	})
	if err != nil {
		return nil, err
	}
	// 解析 XML 响应
	var resp SignDataByP7AttachResponse
	if err = xml.Unmarshal([]byte(httpResp), &resp); err != nil {
		logx.WithContext(ctx).Errorf("SignDataByP7Attach Response: %+v", err)
		return nil, err
	}
	//encodeToString := base64.StdEncoding.EncodeToString([]byte(resp.Body.SignDataByP7AttachResp.SignDataByP7AttachResult))
	signDataResp := resp.Body.SignDataByP7AttachResp.SignDataByP7AttachResult
	return &bjcatypes.SignDataByP7AttachResp{
		P7SignAttach: signDataResp,
	}, err
}

/*
* @Description: PKCS7 签名(带原文)-验证
* @Author: LiuQHui
* @Receiver b
* @Param ctx
* @Param req
* @Date 2024-05-08 22:27:31
 */
func (b *BjCaCryptWebserviceImplClient) VerifySignedDataByP7Attach(ctx context.Context, req bjcatypes.VerifySignedDataByP7AttachReq) (*bjcatypes.VerifySignedDataByP7AttachResp, error) {
	param := VerifySignedDataByP7AttachReq{
		AppName:       b.Config.AppName,
		Pkcs7SignData: req.Pkcs7SignData,
	}
	httpResp, err := b._sendHttp(ctx, ApiVerifySignedDataByP7Attach, param, nil)
	if err != nil {
		return nil, err
	}
	// 解析 XML 响应
	var resp VerifySignedDataByP7AttachResponse
	if err = xml.Unmarshal([]byte(httpResp), &resp); err != nil {
		logx.WithContext(ctx).Errorf("SignDataByP7Attach Response: %+v", err)
		return nil, err
	}

	return &bjcatypes.VerifySignedDataByP7AttachResp{
		VerifyRes: resp.Body.VerifySignedDataByP7AttachResp.VerifySignedDataByP7Attach,
	}, err
}
