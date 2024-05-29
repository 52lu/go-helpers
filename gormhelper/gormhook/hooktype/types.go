package hooktype

import "gorm.io/gorm"

type HookPluginConf struct {
	DB                   *gorm.DB
	FilterDiffColumnList []string
}
