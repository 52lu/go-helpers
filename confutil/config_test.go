package confutil

import (
	"fmt"
	"github.com/52lu/go-helpers/confutil/conftype"
	"testing"
	"time"
)

func initConfig() error {
	// 解析配置
	client, err := NewConfigParseClient(conftype.ConfigParseConf{
		ConfigPaths: []string{"./tmp"},
		ConfigFile:  "local.toml",
		ParseMethod: conftype.ParseMethodTypeViper,
		ApolloConf: &conftype.ApolloConfig{
			Enable:           true,
			ServiceUrl:       "http://xxxx",
			Cluster:          "default",
			AppId:            "appid",
			Namespaces:       []string{"application", "app.json"},
			IsBackupConfig:   true,
			BackupConfigPath: "./tmp",
		},
	})
	if err != nil {
		return err
	}
	err = client.ParseConfig()
	if err != nil {
		return err
	}
	return nil
}

func TestParse(t *testing.T) {
	// 配置初始化
	err := initConfig()
	if err != nil {
		t.Error(err)
		return
	}
	s := GetString("common_conf")
	fmt.Println(s)
	fmt.Println("获取key ENABLE_SWITH_TYPE:", GetInt64("ENABLE_SWITH_TYPE"))
	for true {
		time.Sleep(time.Second)
		fmt.Println("获取key ENABLE_SWITH_TYPE:", GetInt64("ENABLE_SWITH_TYPE"))

		fmt.Println("获取---test_env:", GetStringSlice("test_env"))
		fmt.Println("获取---api.url:", GetString("api.url"))
	}

	fmt.Println("ok")
}
