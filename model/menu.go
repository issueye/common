package model

import (
	"strconv"

	"golang.corp.yxkj.com/guava/common/pkg/utils"
)

type Menu struct {
	Base
	MenuBase
}

type MenuBase struct {
	Name     string `binding:"required" label:"菜单名称" gorm:"column:name;type:nvarchar(200);comment:菜单名称;" json:"name"`              // 菜单名称
	Title    string `binding:"required" label:"菜单标题" gorm:"column:title;type:nvarchar(200);comment:菜单标题;" json:"title"`            // 菜单标题
	Route    string `binding:"required" label:"菜单路由" gorm:"column:route;type:nvarchar(400);comment:菜单路由;" json:"route"`            // 路由
	Icon     string `label:"菜单图标" gorm:"column:icon;type:nvarchar(100);comment:菜单图标;" json:"icon"`                                 // 菜单图标
	Order    int    `label:"排序" gorm:"column:order;type:int;comment:排序;" json:"order"`                                             // 排序
	Auth     int    `label:"菜单权限级别" gorm:"column:auth;type:int;comment:菜单权限级别;" json:"auth"`                                       // 菜单权限级别
	Level    int    `label:"菜单层级级别" gorm:"column:level;type:int;comment:菜单层级级别;" json:"level"`                                     // 菜单层级级别 0 代表第一层
	ParentId string `binding:"required" label:"父级菜单编码" gorm:"column:parent_id;type:nvarchar(100);comment:父级菜单编码;" json:"parentId"` // 父级菜单编码 如果没有则传"0"
	State    uint   `label:"菜单状态" gorm:"column:state;type:int;comment:菜单状态;" json:"state"`                                         // 状态 0 停用 1 启用
}

func (mod *MenuBase) Copy(data *MenuBase) {
	mod.Name = data.Name
	mod.Title = data.Title
	mod.Route = data.Route
	mod.Icon = data.Icon
	mod.Order = data.Order
	mod.Auth = data.Auth
	mod.Level = data.Level
	mod.ParentId = data.ParentId
	mod.State = data.State
}

// TableName
// 表名称
func (Menu) TableName() string {
	return "menu_info"
}

func (Menu) New() *Menu {
	return &Menu{
		Base: Base{
			ID: strconv.FormatInt(utils.GenID(), 10),
		},
	}
}

type GroupMenu struct {
	Base
	GroupMenuBase
}

type GroupMenuBase struct {
	GroupId string `binding:"required" label:"用户组编码" gorm:"column:group_id;primaryKey;type:nvarchar(100);comment:用户组编码;" json:"group_id"` // 用户组编码
	MenuId  string `binding:"required" label:"菜单编码" gorm:"column:menu_id;type:nvarchar(100);comment:菜单编码;" json:"menu_id"`                // 菜单编码
	MenuBase
}

func (mod *GroupMenu) Copy(data *GroupMenuBase) {
	mod.MenuId = data.MenuId
	mod.GroupId = data.GroupId
	mod.Name = data.Name
	mod.Title = data.Title
	mod.Route = data.Route
	mod.Icon = data.Icon
	mod.Order = data.Order
	mod.Auth = data.Auth
	mod.Level = data.Level
	mod.ParentId = data.ParentId
	mod.State = data.State
}

func (mod *GroupMenu) Copy2(data *MenuBase) {
	mod.Name = data.Name
	mod.Title = data.Title
	mod.Route = data.Route
	mod.Icon = data.Icon
	mod.Order = data.Order
	mod.Auth = data.Auth
	mod.Level = data.Level
	mod.ParentId = data.ParentId
	mod.State = data.State
}

// TableName
// 表名称
func (GroupMenu) TableName() string {
	return "group_menu_info"
}

func (GroupMenu) New() *GroupMenu {
	return &GroupMenu{
		Base: Base{
			ID: strconv.FormatInt(utils.GenID(), 10),
		},
	}
}
