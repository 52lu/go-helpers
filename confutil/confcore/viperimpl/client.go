package viperimpl

import (
	"github.com/52lu/go-helpers/confutil/confcore/apolloconf"
	"github.com/52lu/go-helpers/confutil/conftype"
	"github.com/spf13/viper"
	"path/filepath"
	"strings"
	"sync"
)

type viperParseInstance struct {
	conf  conftype.ConfigParseConf
	viper *viper.Viper
}

/*
* @Description: 配置实例
* @Author: LiuQHui
* @Date 2024-05-31 18:57:39
 */
func NewViperConfInstance(cf conftype.ConfigParseConf) *viperParseInstance {
	viperInstance := viper.New()
	/*
		  当使用SetConfigFile时，不需要再使用 SetConfigName或 AddConfigPath。
			SetConfigFile：需要完整的文件路径，包括文件名和扩展名。
	*/
	if filepath.IsAbs(cf.ConfigFile) {
		// 设置配置文件;SetConfigFil 需要完整的文件路径，包括文件名和扩展名。
		viperInstance.SetConfigFile(cf.ConfigFile)
	} else {
		// 当不是绝对路径时，使用SetConfigName和SetConfigType
		// 提取文件名和扩展名
		filename := filepath.Base(cf.ConfigFile) // 提取文件名（包括扩展名）
		extension := filepath.Ext(cf.ConfigFile) // 提取扩展名
		// 设置配置文件名，没有后缀
		viperInstance.SetConfigName(filename[:len(filename)-len(extension)])
		// 设置读取文件格式
		viperInstance.SetConfigType(strings.ReplaceAll(extension, ".", ""))
		// 设置配置文件目录(可以设置多个,优先级根据添加顺序来)
		if len(cf.ConfigPaths) > 0 {
			for _, path := range cf.ConfigPaths {
				viperInstance.AddConfigPath(path)
			}
		}
	}
	return &viperParseInstance{
		conf:  cf,
		viper: viperInstance,
	}
}

/*
* @Description: 配置解析
* @Author: LiuQHui
* @Receiver v
* @Return error
* @Date 2024-06-04 12:28:59
 */
func (v *viperParseInstance) Parse() error {
	err := v.viper.ReadInConfig()
	if err != nil {
		return err
	}
	// 是否开启apollo
	if v.conf.ApolloConf != nil && v.conf.ApolloConf.Enable {
		// 合并apollo配置
		err = v.mergeApollo()
	}
	return err
}

/*
* @Description: 更新配置
* @Author: LiuQHui
* @Receiver v
* @Param confMap
* @Return error
* @Date 2024-06-05 14:36:21
 */
func (v *viperParseInstance) UpdateConf(confMap map[string]interface{}) error {
	v._updateViper(confMap)
	return nil
}

/*
* @Description: heb
* @Author: LiuQHui
* @Receiver v
* @Return error
* @Date 2024-06-05 14:28:13
 */
func (v *viperParseInstance) mergeApollo() error {
	client := apolloconf.NewApolloConfClient(v.conf.ApolloConf)
	client.SetConfigChangeNotifyImpl(v)
	err := client.Start()
	if err != nil {
		return err
	}
	for _, namespace := range v.conf.ApolloConf.Namespaces {
		confMap, _ := client.GetConfMapByNamespace(namespace)
		if len(confMap) == 0 {
			continue
		}
		v._updateViper(confMap)
	}
	return nil
}

func (v *viperParseInstance) _updateViper(confMap map[string]interface{}) {
	var lock sync.Mutex
	// 合并到viper
	for key, val := range confMap {
		lock.Lock()
		v.viper.Set(key, val)
		lock.Unlock()
	}
}
