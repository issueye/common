package model

import (
	"strconv"

	"golang.corp.yxkj.com/guava/common/pkg/utils"
)

// User
// 用户信息
type UserInfo struct {
	Base
	UserBase
}

type UserBase struct {
	GroupId  string `gorm:"column:group_id;type:nvarchar(100)" json:"groupId"` // 组编码
	Account  string `gorm:"column:account;type:nvarchar(50)" json:"account"`   // uid 登录名
	Name     string `gorm:"column:name;type:nvarchar(50)" json:"name"`         // 用户姓名
	Password string `gorm:"column:password;type:nvarchar(50)" json:"password"` // 密码
	Mark     string `gorm:"column:mark;type:nvarchar(500)" json:"mark"`        // 备注
	State    uint   `gorm:"column:state;type:int" json:"state"`                // 状态 0 停用 1 启用
	Sys      uint   `gorm:"column:sys;type:int" json:"sys"`                    // 系统生成 0 否 1 是
}

func (mod *UserInfo) Copy(data *UserBase) {
	mod.Name = data.Name
	mod.GroupId = data.GroupId
	mod.Account = data.Account
	mod.Password = data.Password
	mod.Mark = data.Mark
	mod.State = data.State
	mod.Sys = data.Sys
}

func (UserInfo) TableName() string {
	return "user_info"
}

func (UserInfo) New() *UserInfo {
	return &UserInfo{
		Base: Base{
			ID: strconv.FormatInt(utils.GenID(), 10),
		},
	}
}

// 用户组信息
type UserGroupInfo struct {
	Base
	UserGroupBase
}

type UserGroupBase struct {
	Name  string `gorm:"column:name;type:nvarchar(300);comment:组名称;" json:"name"` // 组名称
	Mark  string `gorm:"column:mark;type:nvarchar(2000);comment:备注;" json:"mark"` // 备注
	State uint   `gorm:"column:state;type:int" json:"state"`                      // 状态 0 停用 1 启用
	Sys   uint   `gorm:"column:sys;type:int" json:"sys"`                          // 系统生成 0 否 1 是
}

func (mod *UserGroupInfo) Copy(data *UserGroupBase) {
	mod.Name = data.Name
	mod.Mark = data.Mark
	mod.State = data.State
	mod.Sys = data.Sys
}

func (UserGroupInfo) TableName() string {
	return "user_group_info"
}

func (UserGroupInfo) New() *UserGroupInfo {
	return &UserGroupInfo{
		Base: Base{
			ID: strconv.FormatInt(utils.GenID(), 10),
		},
	}
}
