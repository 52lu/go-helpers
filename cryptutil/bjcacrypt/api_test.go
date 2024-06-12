package bjcacrypt

import (
	"context"
	"fmt"
	"gitlab.dev.olanyun.com/olanyun/saas-utils/cryptoutil/bjcacrypt/bjcatypes"
	"gitlab.dev.olanyun.com/olanyun/saas-utils/jsonutil"
	"testing"
)

var (
	ctx          = context.TODO()
	newclient, _ = NewBjCaCryptImplClient(ctx, &bjcatypes.CaConfig{
		Url:        "",
		AppName:    "SVSDefault",
		BJCACertSn: "2136e7af168f2898",
		ImplType:   ImplTypeWebservice,
		IsTest:     true,
		//RedisConfig: &bjcatypes.RedisConfig{
		//	Host: "127.0.0.1:6379",
		//	Pass: "",
		//},
	})
)

/*
* @Description: 获取证书
* @Author: LiuQHui
* @Param t
* @Date 2023-12-01 17:41:16
 */
func TestGetServerCertificate(t *testing.T) {
	certificate, err := newclient.GetCertBySn(ctx)

	fmt.Println(err)
	fmt.Println(certificate)
}

/*
* @Description: 数字信封加密
* @Author: LiuQHui
* @Param t
* @Date 2023-12-01 17:41:22
 */
func TestEncodeEnvelopedData(t *testing.T) {
	tmp := struct {
		Name  string `json:"name"`
		Age   int64  `json:"age"`
		Phone string `json:"phone"`
	}{
		Name:  "张三",
		Age:   18,
		Phone: "176000202001",
	}

	toString, _ := jsonutil.Json().MarshalToString(tmp)

	resp, err := newclient.EncodeEnvelopedData(ctx, bjcatypes.EncodeEnvelopedDataReq{
		InData: toString,
	})
	fmt.Println(err)
	fmt.Println(resp)
}

/*
* @Description: 数字信封解密
* @Author: LiuQHui
* @Param t
* @Date 2023-12-01 17:43:57
 */
func TestDecodeEnvelopedData(t *testing.T) {
	resp, err := newclient.DecodeEnvelopedData(ctx, bjcatypes.DecodeEnvelopedDataReq{
		InData: "MIIBPgYKKoEcz1UGAQQCA6CCAS4wggEqAgEAMYHmMIHjAgEAMFIwRDELMAkGA1UEBhMCQ04xDTALBgNVBAoMBEJKQ0ExDTALBgNVBAsMBEJKQ0ExFzAVBgNVBAMMDkJlaWppbmcgU00yIENBAgoaEAAAAAAQs9pEMA0GCSqBHM9VAYItAwUABHsweQIhANZVY4CQN8poIxhGLP9vGSzTDzkXnuQyDw28ssTBZvhQAiBzgQO8yy7Wit8WCvwe7j3QJjf/NoifmJ668jzOB62IiQQgz1TR5YMcLbUUIPuvPm8z3oMIMbZ9JE3ofHI0fHMDz8AEEFDI2VYbsDFDda7T0xG5I/swPAYKKoEcz1UGAQQCATAcBggqgRzPVQFoAgQQgackp9UTu3L9KYUNjkulOIAQtFPhGFxDQJaXaCtW/kF9PQ==",
	})
	fmt.Println(err)
	fmt.Println(resp)
}

/*
* @Description: PKCS7签名(带原文)-生成签名
* @Author: LiuQHui
* @Param t
* @Date 2023-12-01 17:45:33
 */
func TestSignDataByP7Attach(t *testing.T) {
	resp, err := newclient.SignDataByP7Attach(ctx, bjcatypes.SignDataByP7AttachReq{
		InData: "MIIEOwYKKoEcz1UGAQQCA6CCBCswggQnAgEAMYHwMIHtAgEAMFwwUDELMAkGA1UEBhMCQ04xDTALBgNVBAoMBE5IU0ExMjAwBgNVBAMMKVByaXZhdGUgQ2VydGlmaWNhdGUgQXV0aG9yaXR5IE9mIE5IU0EgU00yAgghNuevFo8omDANBgkqgRzPVQGCLQMFAAR7MHkCIDOUpBydMimvxuF8YsMD4IKA1Av6RRnoSk1daeoRD7kRAiEAwcccvMAkYA2XZG3yuu00E5/5raXlxCShh4aXIZ2d19MEIH8B9a308hbpwRigpop1+nBegz+buTsKLDPC6RRhpCRXBBB9h9fFu8d2iK54ZELD7vGjMIIDLQYKKoEcz1UGAQQCATAbBgcqgRzPVQFoBBCfE5BxZn/Dixpp/XFeFleLgIIDAGCgwVPOuSIgwkBHCieULpRZ3WAAC3q+2AbcRfxtHkhUidT1b3QdGQojXwwJkhFR7aeKfTj0zCH9zHdGJTERFJ+DuF1tqNHYeZhcwn9iZLxWY/KO+Y4dCqf7DjuwpoPm2ct3zTqeKirCk8BLjCykP90Q2FDuvMnqZv5e4JfFXmdK1FyFniotdFQwvZzJjM5Omw132MbxGrV058BNc/xZceVCsPl7ke93x4SyT5MVSH76WTDryF1mOmrlAIoEj5HgIoBBNb8UgSv9GJhlWfLoufe+Yf6yPdeyf0cdXbzfIgOBleMkW9AWB8Uy/imwrwQAJs2JeNJWMPXM+U+FWhiFgBkTewX4yWgNmZkoYWToAcMedlfHl7Kja6CwbfgbcPzRPmRGCe73iWmcpdWrsxrAOBYwg5xhDq4970bM2LFMxRHCdSoXxUb/8w756J4V2eCZlqtpGA28cdVwNDtAysVgJKw99ebBT7+MpTjcW6yzdeYFBC+iXjwG9eybbVElymbPd1uW9Fkjal4bN2E+lGbGsnnAjNnRDIzUuZJh9c6yJaoXZvxObLOJlakMVt3W0kFoYpjJgNgTIX9pTMmwfcyuGX8FX8ROrpNXWNLHaQV06bdYH/Ge0M2+a5OuKpZxi9ggSRMIidq20gsy6f12eIXNFPs4caBeZo5yLvEKSlHM9BbRKNQNiGid8MH9DYmlac4f+c8jRxZjD81WgNBC4tsbdlu0lsETUpQQc0F3vYNtTQiEdULdi3WzXKcTbRHdkWAkWzIotn+EKT5EnCClI5dG2vrGFdx68o6TT6MRQVdMHCC3N4Dky0sTyYIXAgXRZXIpF2VCwGPHwxNL7TG2ttk0ZBMnMDCeV07U7apJwFnJXCcN8R7kRx7eXGQZEQrlzqdTnIpoGpnufYTqSVsu1OHAeGL6G9Irundi5Cg+ZP0BkWMxxh5hr8oQKhktXdFLXwW2KId6BG/44ziEATtWXIobAKOZJuNHd6HMT9a1vioOS4xtg5AEOpFaO0LK5/K7Kw7sCg==",
	})
	fmt.Println(err)
	fmt.Println(resp)
}

/*
* @Description: PKCS7签名(带原文)-验证签名
* @Author: LiuQHui
* @Param t
* @Date 2023-12-01 17:47:10
 */
func TestVerifySignedDataByP7Attach(t *testing.T) {
	resp, err := newclient.VerifySignedDataByP7Attach(ctx, bjcatypes.VerifySignedDataByP7AttachReq{
		Pkcs7SignData: "MIIHEwYKKoEcz1UGAQQCAqCCBwMwggb/AgEBMQ4wDAYIKoEcz1UBgxEFADCCAcQGCiqBHM9VBgEEAgGgggG0BIIBsE1JSUJQZ1lLS29FY3oxVUdBUVFDQTZDQ0FTNHdnZ0VxQWdFQU1ZSG1NSUhqQWdFQU1GSXdSREVMTUFrR0ExVUVCaE1DUTA0eERUQUxCZ05WQkFvTUJFSktRMEV4RFRBTEJnTlZCQXNNQkVKS1EwRXhGekFWQmdOVkJBTU1Ea0psYVdwcGJtY2dVMDB5SUVOQkFnb2FFQUFBQUFBUXM5cEVNQTBHQ1NxQkhNOVZBWUl0QXdVQUJIc3dlUUloQU5aVlk0Q1FOOHBvSXhoR0xQOXZHU3pURHprWG51UXlEdzI4c3NUQlp2aFFBaUJ6Z1FPOHl5N1dpdDhXQ3Z3ZTdqM1FKamYvTm9pZm1KNjY4anpPQjYySWlRUWd6MVRSNVlNY0xiVVVJUHV2UG04ejNvTUlNYlo5SkUzb2ZISTBmSE1EejhBRUVGREkyVllic0RGRGRhN1QweEc1SS9zd1BBWUtLb0VjejFVR0FRUUNBVEFjQmdncWdSelBWUUZvQWdRUWdhY2twOVVUdTNMOUtZVU5qa3VsT0lBUXRGUGhHRnhEUUphWGFDdFcva0Y5UFE9PaCCBFwwggRYMIID/6ADAgECAgoaEAAAAAAQs9pEMAoGCCqBHM9VAYN1MEQxCzAJBgNVBAYTAkNOMQ0wCwYDVQQKDARCSkNBMQ0wCwYDVQQLDARCSkNBMRcwFQYDVQQDDA5CZWlqaW5nIFNNMiBDQTAeFw0yMzAyMDcxNjAwMDBaFw0yNDAyMDgxNTU5NTlaMIGLMRAwDgYDVQQpDAcxMDAwMDAwMRAwDgYDVQQDDAdzbTJ0ZXN0MR4wHAYDVQQLDBXkuK3lm73pk4Hot6/mgLvlhazlj7gxHjAcBgNVBAoMFeS4reWbvemTgei3r+aAu+WFrOWPuDELMAkGA1UEBwwCIiIxCzAJBgNVBAgMAiIiMQswCQYDVQQGDAJDTjBZMBMGByqGSM49AgEGCCqBHM9VAYItA0IABKBln3wEzGYSvfqikJaYdyYfKknNpd9ccvaKO7M76CXnlwAERWPYcJperHpoJMIH37u/uQipEwO8p2qYvlSpteijggKPMIICizAfBgNVHSMEGDAWgBQf5s/Uj8UiKpdKKYoV5xbJkjTEtjAdBgNVHQ4EFgQUCHQmjkMWq816RYlZ2qROocp8tbIwDgYDVR0PAQH/BAQDAgbAMIGjBgNVHR8EgZswgZgwYKBeoFykWjBYMQswCQYDVQQGEwJDTjENMAsGA1UECgwEQkpDQTENMAsGA1UECwwEQkpDQTEXMBUGA1UEAwwOQmVpamluZyBTTTIgQ0ExEjAQBgNVBAMTCWNhMjFjcmwzMDA0oDKgMIYuaHR0cDovL3Rlc3QuYmpjYS5vcmcuY246ODAwMy9jcmwvY2EyMWNybDMwLmNybDAZBgoqgRyG7zICAQEBBAsMCUpKMTAwMDAwMDBgBggrBgEFBQcBAQRUMFIwIwYIKwYBBQUHMAGGF09DU1A6Ly9vY3NwLmJqY2Eub3JnLmNuMCsGCCsGAQUFBzAChh9odHRwOi8vY3JsLmJqY2Eub3JnLmNuL2NhaXNzdWVyMEAGA1UdIAQ5MDcwNQYJKoEchu8yAgIBMCgwJgYIKwYBBQUHAgEWGmh0dHA6Ly93d3cuYmpjYS5vcmcuY24vY3BzMBEGCWCGSAGG+EIBAQQEAwIA/zAXBgoqgRyG7zICAQEIBAkMBzEwMDAwMDAwGQYKKoEchu8yAgECAgQLDAlKSjEwMDAwMDAwHwYKKoEchu8yAgEBDgQRDA85OTgwMDAxMDA3MDc2NTUwGQYKKoEchu8yAgEBBAQLDAlKSjEwMDAwMDAwJAYKKoEchu8yAgEBFwQWDBQ2MjFAMjE1MDA5SkowMTAwMDAwMDAVBggqgRzQFAQBBAQJDAcxMDAwMDAwMBQGCiqBHIbvMgIBAR4EBgwEMTA1MDAKBggqgRzPVQGDdQNHADBEAiAoIYd8cK1Gp3rT2akPNP3l1ExubImmcJxXtKgdm9Z/DgIgHv/P8EDviAEgONdWZZCqw1+edGaXGpnvrNDsm5SE86cxgcEwgb4CAQEwUjBEMQswCQYDVQQGEwJDTjENMAsGA1UECgwEQkpDQTENMAsGA1UECwwEQkpDQTEXMBUGA1UEAwwOQmVpamluZyBTTTIgQ0ECChoQAAAAABCz2kQwDAYIKoEcz1UBgxEFADANBgkqgRzPVQGCLQEFAARIMEYCIQCHiMaP+YCzHYQ5xANsNilNIWBCM/qtzsHRdNwaPuMRTQIhAOExliqtG4NGPjl9AtJqTxKlYeYzwOLyb3gjERApSkTz",
	})
	fmt.Println(err)
	fmt.Println(resp)
}
