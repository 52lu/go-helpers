package confutil

import (
	"fmt"
	"github.com/52lu/go-helpers/confutil/confcore/viperimpl"
	"github.com/52lu/go-helpers/confutil/conftype"
)

type configParseClient struct {
	impl conftype.ConfigParseInterface
}

var _client *configParseClient

/*
* @Description: 实例化配置客户端
* @Author: LiuQHui
* @Param conf
* @Return *configParseClient
* @Return error
* @Date 2024-06-04 14:09:50
 */
func NewConfigParseClient(conf conftype.ConfigParseConf) (*configParseClient, error) {
	if conf.ParseMethod != conftype.ParseMethodTypeViper {
		return nil, fmt.Errorf("conf.ParseMethod %v 未实现", conf.ParseMethod)
	}
	// 实例化实现
	_client = &configParseClient{
		impl: viperimpl.NewViperConfInstance(conf),
	}
	return _client, nil
}

/*
* @Description: 配置解析
* @Author: LiuQHui
* @Receiver c
* @Return error
* @Date 2024-06-04 14:16:23
 */
func (c *configParseClient) ParseConfig() error {
	return c.impl.Parse()
}
