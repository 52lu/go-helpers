package viperimpl

import (
	"52lu/go-helpers/confutil/conftype"
	"fmt"
	"github.com/spf13/viper"
	"path/filepath"
	"strings"
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
		fmt.Println("fileName:", filename[:len(filename)-len(extension)])
		fmt.Println("extension:", strings.ReplaceAll(extension, ".", ""))
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
	return nil
}
