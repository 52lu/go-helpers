package apolloconf

import (
	"fmt"
	"github.com/52lu/go-helpers/confutil/conftype"
	"github.com/52lu/go-helpers/jsonutil"
	"github.com/apolloconfig/agollo/v4"
	"github.com/apolloconfig/agollo/v4/component/log"
	"github.com/apolloconfig/agollo/v4/env/config"
	"github.com/apolloconfig/agollo/v4/storage"
	"github.com/panjf2000/ants/v2"
	"strings"
	"sync"
)

type apolloConfClient struct {
	conf           *conftype.ApolloConfig
	apollo         agollo.Client
	logger         log.LoggerInterface
	confNotifyImpl ConfigChangeNotifyInterface
	confListener   storage.ChangeListener
}

type ConfigChangeNotifyInterface interface {
	UpdateConf(confMap map[string]interface{}) error
}

/*
* @Description: 实例化客户端
* @Author: LiuQHui
* @Param conf
* @Return *apolloConfClient
* @Date 2024-06-05 11:46:01
 */
func NewApolloConfClient(conf *conftype.ApolloConfig) *apolloConfClient {
	return &apolloConfClient{conf: conf}
}

/*
* @Description: 设置日志插件
* @Author: LiuQHui
* @Receiver a
* @Param logger
* @Date 2024-06-05 11:49:25
 */
func (a *apolloConfClient) SetLogger(logger log.LoggerInterface) {
	a.logger = logger
}

/*
* @Description: 设置监听器
* @Author: LiuQHui
* @Receiver a
* @Param logger
* @Date 2024-06-05 11:49:25
 */
func (a *apolloConfClient) SetListener(listener storage.ChangeListener) {
	a.confListener = listener
}

/*
* @Description: 设置配置变更回调
* @Author: LiuQHui
* @Receiver a
* @Param impl
* @Date 2024-06-05 14:39:19
 */
func (a *apolloConfClient) SetConfigChangeNotifyImpl(impl ConfigChangeNotifyInterface) {
	a.confNotifyImpl = impl
}

/*
* @Description: 启动服务
* @Author: LiuQHui
* @Receiver a
* @Return error
* @Date 2024-06-05 11:47:29
 */
func (a *apolloConfClient) Start() error {
	// 初始化apollo
	c := &config.AppConfig{
		AppID:             a.conf.AppId,
		Cluster:           a.conf.Cluster,
		NamespaceName:     strings.Join(a.conf.Namespaces, ","),
		IP:                a.conf.ServiceUrl,
		IsBackupConfig:    a.conf.IsBackupConfig,
		BackupConfigPath:  a.conf.BackupConfigPath,
		Secret:            a.conf.Secret,
		Label:             a.conf.AppId,
		SyncServerTimeout: a.conf.SyncServerTimeout,
	}
	// 设置日志实现
	if a.logger != nil {
		agollo.SetLogger(a.logger)
	}
	var err error
	// 启动apollo
	a.apollo, err = agollo.StartWithConfig(func() (*config.AppConfig, error) {
		return c, nil
	})
	// 启动监听器
	a.enableListener()

	return err
}

/*
* @Description: 开启监听
* @Author: LiuQHui
* @Receiver a
* @Date 2024-06-05 15:45:08
 */
func (a *apolloConfClient) enableListener() {
	// 设置监听器
	if a.confListener != nil {
		a.SetListener(a.confListener)
	} else {
		// 默认监视器
		a.SetListener(NewApolloChangeListener(a.confNotifyImpl))
	}
	// 开启
	_ = ants.Submit(func() {
		var wg sync.WaitGroup
		wg.Add(1)
		a.apollo.AddChangeListener(a.confListener)
		wg.Wait()
	})
}

/*
* @Description: 通过命名空间获取配置
* @Author: LiuQHui
* @Receiver a
* @Param namespace
* @Return map[string]interface{}
* @Return error
* @Date 2024-06-05 14:10:04
 */
func (a *apolloConfClient) GetConfMapByNamespace(namespace string) (map[string]interface{}, error) {
	var namespaceConMap = make(map[string]interface{})
	a.apollo.GetConfigCache(namespace).Range(func(key, value interface{}) bool {
		if key == "content" {
			if valStr, ok := value.(string); ok {
				_ = jsonutil.Json.UnmarshalFromString(valStr, &namespaceConMap)
			}
			return true
		}
		namespaceConMap[fmt.Sprintf("%v", key)] = value
		return true
	})
	return namespaceConMap, nil
}
