package httpimpl

import (
	"context"
	"fmt"
	"github.com/cockroachdb/errors"
	"github.com/zeromicro/go-zero/core/logx"
	"github.com/zeromicro/go-zero/core/stores/redis"
	"gitlab.dev.olanyun.com/olanyun/saas-utils/cryptoutil/bjcacrypt/bjcatypes"
	"gitlab.dev.olanyun.com/olanyun/saas-utils/httputil"
	"gitlab.dev.olanyun.com/olanyun/saas-utils/jsonutil"
)

type BjCaCryptHttpImplClient struct {
	Config *bjcatypes.CaConfig
	rd     *redis.Redis
}

var (
	jsonhelper = jsonutil.Json()
)

/*
* @Description: 实例化客户端
* @Author: LiuQHui
* @Param ctx
* @Param conf
* @Date 2023-12-01 14:28:03
 */
func NewCaCryptHttpImplClient(ctx context.Context, conf *bjcatypes.CaConfig) (*BjCaCryptHttpImplClient, error) {
	// 实例化redis
	var rd *redis.Redis
	if conf.RedisConfig != nil {
		rd = redis.New(conf.RedisConfig.Host, func(r *redis.Redis) {
			r.Type = redis.NodeType
			r.Pass = conf.RedisConfig.Pass
		})
	}
	client := &BjCaCryptHttpImplClient{
		Config: conf,
		rd:     rd,
	}
	if conf.BJCABase64Cert == "" {
		// 获取证书
		certificateResp, err := client.GetCertBySn(ctx)
		if err != nil {
			return nil, err
		}
		client.Config.BJCABase64Cert = certificateResp.Base64Cert
	}

	return client, nil

}

/*
* @Description: 发起请求
* @Author: LiuQHui
* @Receiver c
* @Param ctx
* @Param apiName
* @Param param
* @Date 2023-12-01 14:39:50
 */
func (c *BjCaCryptHttpImplClient) sendHttp(ctx context.Context, apiName string, param interface{}) (string, error) {
	url := fmt.Sprintf("%s%s", c.Config.Url, apiName)
	toString, err := jsonhelper.MarshalToString(param)
	if err != nil {
		return "", err
	}
	respStr, err := httputil.SendPostJsonRequest(ctx, url, toString, nil)
	logx.WithContext(ctx).Infow("数字认证请求记录"+url,
		logx.LogField{Key: "请求地址", Value: url},
		logx.LogField{Key: "入参", Value: param},
		logx.LogField{Key: "返回", Value: respStr},
	)
	var responseResult bjcatypes.CommonResp
	err = jsonhelper.UnmarshalFromString(respStr, &responseResult)
	if err != nil {
		logx.WithContext(ctx).Errorf("返回结果解析失败 %+v %v", err, respStr)
		return "", err
	}
	if responseResult.Status != ResponseSuccessCode || responseResult.MsgInfo != ResponseSuccessMsg {
		paramStr, _ := jsonhelper.MarshalToString(param)
		var errMsg string
		// 异常处理
		marshalToString, err := jsonhelper.MarshalToString(responseResult.Body)
		if err == nil {
			var errResp ErrorRespProto
			if err = jsonhelper.UnmarshalFromString(marshalToString, &errResp); err == nil {
				errMsg = fmt.Sprintf("%v|%v", errResp.ErrCode, errResp.DetailMsg)
			}
		}
		logx.WithContext(ctx).Errorf("数字证书响应失败: %s \n url: %s \n 入参:%s \n 返回: %s",
			errMsg, url, paramStr, respStr)
		return "", errors.New(errMsg)
	}
	return jsonhelper.MarshalToString(responseResult.Body)
}

type ErrorRespProto struct {
	DetailMsg string `json:"detailMsg"`
	ErrCode   int    `json:"errCode"`
}
