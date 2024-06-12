package websvcimpl

import (
	"bytes"
	"context"
	"encoding/xml"
	"fmt"
	"github.com/zeromicro/go-zero/core/logx"
	"github.com/zeromicro/go-zero/core/stores/redis"
	"gitlab.dev.olanyun.com/olanyun/saas-utils/cryptoutil/bjcacrypt/bjcatypes"
	"gitlab.dev.olanyun.com/olanyun/saas-utils/errutil"
	"gitlab.dev.olanyun.com/olanyun/saas-utils/jsonutil"
	"html"
	"io"
	"io/ioutil"
	"net/http"
	"regexp"
)

type BjCaCryptWebserviceImplClient struct {
	Config *bjcatypes.CaConfig
	rd     *redis.Redis
}

var (
	jsonhelper = jsonutil.Json()
)

/*
* @Description: webservice client
* @Author: LiuQHui
* @Param ctx
* @Param conf
* @Date 2024-05-08 19:20:03
 */
func NewCaCryptWebserviceImplClient(ctx context.Context, conf *bjcatypes.CaConfig) (*BjCaCryptWebserviceImplClient, error) {
	// 实例化redis
	var rd *redis.Redis
	if conf.RedisConfig != nil {
		rd = redis.New(conf.RedisConfig.Host, func(r *redis.Redis) {
			r.Type = redis.NodeType
			r.Pass = conf.RedisConfig.Pass
		})
	}
	client := &BjCaCryptWebserviceImplClient{
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
* @Description: 执行网络请求
* @Author: LiuQHui
* @Receiver b
* @Param ctx
* @Param body
* @Date 2024-05-16 22:48:19
 */
func (b *BjCaCryptWebserviceImplClient) _execSendHttp(ctx context.Context, body string) (string, error) {
	//req, err := http.NewRequest("POST", b.Config.Url, io.NopCloser(strings.NewReader(body)))
	req, err := http.NewRequest("POST", b.Config.Url, io.NopCloser(bytes.NewReader([]byte(body))))
	if err != nil {
		return "", err
	}
	req.Header.Add("Content-Type", "text/xml")
	// 发起请求数据
	client := http.Client{}
	response, err := client.Do(req)
	if err != nil {
		return "", err
	}
	// 解析返回数据
	if response.StatusCode != 200 {
		return "", errutil.ThrowErrorMsg("网络请求失败")
	}
	// 读取后需要关闭response.Body
	defer func() {
		response.Body.Close()
	}()
	// 读取response.Body
	res, err := ioutil.ReadAll(response.Body)
	// 实体符号反转义
	result := html.UnescapeString(string(res))
	return result, err
}

/*
* @Description: 发送请求
* @Author: LiuQHui
* @Receiver b
* @Param ctx
* @Param apiName
* @Param proto
* @Date 2024-05-08 22:00:41
 */
func (b *BjCaCryptWebserviceImplClient) _sendHttp(ctx context.Context, apiName string,
	proto interface{}, byteDataMap map[string][]byte) (string, error) {
	// 参数转成xml
	protoXml, err := xml.Marshal(proto)
	if err != nil {
		return "", errutil.ThrowError(err)
	}
	format := `<soapenv:Envelope xmlns:soapenv="http://schemas.xmlsoap.org/soap/envelope/" xmlns:tem="http://tempuri.org/">
   <soapenv:Header/>
   <soapenv:Body>
		%v
   </soapenv:Body>
</soapenv:Envelope>
`
	tmpBody := string(protoXml)
	if len(byteDataMap) > 0 {
		for key, val := range byteDataMap {
			// 编译正则表达式
			regexpPattern := fmt.Sprintf("<tem:%s>.*?</tem:%s>", key, key)
			logx.WithContext(ctx).Infof("正则匹配模式: %s", regexpPattern)
			re := regexp.MustCompile(regexpPattern)
			// 替换
			tmpBody = re.ReplaceAllString(tmpBody, fmt.Sprintf("<tem:%s>%v</tem:%s>", key, val, key))
		}
	}
	body := fmt.Sprintf(format, tmpBody)
	// 发起请求
	var result string
	if b.Config.IsTest {
		result = getTestResponse(apiName)
	} else {
		result, err = b._execSendHttp(ctx, body)
	}
	logx.WithContext(ctx).Infow("bjca(webservice)请求记录-"+apiName, []logx.LogField{
		{Key: "IsTest", Value: b.Config.IsTest},
		{Key: "入参", Value: body},
		{Key: "响应", Value: result},
		{Key: "请求地址", Value: b.Config.Url},
	}...)
	return result, err
}
