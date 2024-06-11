package apolloconf

import (
	"fmt"
	"github.com/52lu/go-helpers/jsonutil"
	"github.com/apolloconfig/agollo/v4/storage"
	"sync"
)

type apolloChangeListener struct {
	notify ConfigChangeNotifyInterface
	lock   sync.Mutex
}

func NewApolloChangeListener(cfNotify ConfigChangeNotifyInterface) *apolloChangeListener {
	return &apolloChangeListener{
		notify: cfNotify,
	}
}

/*
* @Description: 变更监控
* @Author: LiuQHui
* @Receiver a
* @Param event
* @Date 2024-06-05 14:50:51
 */
func (a *apolloChangeListener) OnChange(event *storage.ChangeEvent) {
	if a.notify == nil {
		return
	}
	a.lock.Lock()
	defer a.lock.Unlock()
	var confMap = make(map[string]interface{})
	for key, value := range event.Changes {
		if key == "content" {
			if valStr, ok := value.NewValue.(string); ok {
				_ = jsonutil.Json.UnmarshalFromString(valStr, &confMap)
			}
		} else {
			confMap[key] = value.NewValue
		}
	}
	err := a.notify.UpdateConf(confMap)
	if err != nil {
		fmt.Println("OnChange notify.UpdateConf error: ", err)
	}
}

/*
* @Description: 最新变更
* @Author: LiuQHui
* @Receiver a
* @Param event
* @Date 2024-06-05 14:51:08
 */
func (a *apolloChangeListener) OnNewestChange(event *storage.FullChangeEvent) {
	if a.notify == nil {
		return
	}
	a.lock.Lock()
	defer a.lock.Unlock()
	var confMap = make(map[string]interface{})
	for key, value := range event.Changes {
		confMap[key] = value
	}
	err := a.notify.UpdateConf(confMap)
	if err != nil {
		fmt.Println("OnNewestChange notify.UpdateConf error: ", err)
	}
}
