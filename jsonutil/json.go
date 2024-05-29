package jsonutil

import (
	jsoniter "github.com/json-iterator/go"
	"github.com/json-iterator/go/extra"
	"sync"
)

var (
	onceInit sync.Once
	Json     = jsoniter.ConfigCompatibleWithStandardLibrary
)

func init() {
	onceInit.Do(func() {
		// 自适应类型
		extra.RegisterFuzzyDecoders()
	})
}
