package bjcacrypt

import (
	"context"
	"gitlab.dev.olanyun.com/olanyun/saas-utils/cryptoutil/bjcacrypt/bjcatypes"
	"gitlab.dev.olanyun.com/olanyun/saas-utils/cryptoutil/bjcacrypt/httpimpl"
	"gitlab.dev.olanyun.com/olanyun/saas-utils/cryptoutil/bjcacrypt/websvcimpl"
)

type bjcaCryptClient struct {
	ctx context.Context
}

const (
	ImplTypeHttp       = "http"
	ImplTypeWebservice = "webservice"
)

/*
* @Description: 实例化客户端
* @Author: LiuQHui
* @Param ctx
* @Param conf
* @Date 2024-05-08 19:21:36
 */
func NewBjCaCryptImplClient(ctx context.Context, conf *bjcatypes.CaConfig) (BJCACryptInterface, error) {
	if conf.ImplType == ImplTypeWebservice {
		// 实现webservice
		return websvcimpl.NewCaCryptWebserviceImplClient(ctx, conf)
	}
	return httpimpl.NewCaCryptHttpImplClient(ctx, conf)
}
