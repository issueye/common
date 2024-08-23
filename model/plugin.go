package model

import (
	"strconv"

	"github.com/issueye/common/pkg/utils"
)

type PluginInfo struct {
	Base
	PluginBase
}

type PluginBase struct {
	Name    string `binding:"required" label:"名称" gorm:"column:name;size:300;comment:名称;" json:"name"`       // 名称
	Path    string `binding:"required" label:"路径" gorm:"column:path;size:300;comment:路径;" json:"path"`       // 路径
	Version int    `binding:"required" label:"版本" gorm:"column:version;size:300;comment:版本;" json:"version"` // 版本
	Key     string `binding:"required" label:"键" gorm:"column:key;size:300;comment:键;" json:"key"`           // 键
	Value   string `binding:"required" label:"值" gorm:"column:value;size:300;comment:值;" json:"value"`       // 值
}

// TableName
// 表名称
func (PluginInfo) TableName() string {
	return "plugin_info"
}

func (PluginInfo) New() *PluginInfo {
	return &PluginInfo{
		Base: Base{
			ID: strconv.FormatInt(utils.GenID(), 10),
		},
	}
}
