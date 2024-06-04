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
		//ConfigFile:  "/Users/hui/ProjectSpace/GoItem/go-helpers/confutil/tmp/local.toml",
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
	fmt.Println("获取key app_name:", GetString("app_name"))
	fmt.Println("获取整个节点:", Get("apollo"))
	fmt.Println("获取mq节点下的具体key:", Get("mq.access_key_id"))
	databaseList := Get("database")
	fmt.Println(databaseList)
}
