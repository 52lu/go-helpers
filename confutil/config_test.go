package confutil

import (
	"52lu/go-helpers/confutil/conftype"
	"fmt"
	"testing"
)

func TestParse(t *testing.T) {
	// 解析配置
	client, err := NewConfigParseClient(conftype.ConfigParseConf{
		ConfigPaths: []string{"./tmp"},
		ConfigFile:  "local.toml",
		ParseMethod: conftype.ParseMethodTypeViper,
		ApolloConf:  nil,
	})
	if err != nil {
		t.Error(err)
		return
	}
	err = client.ParseConfig()
	if err != nil {
		t.Error(err)
		return
	}

	// 读取配置
	getString := GetString("app_name")
	fmt.Println(getString)

	apolloConf := Get("apollo")
	fmt.Println(apolloConf)
}
