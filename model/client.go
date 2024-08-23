package model

import (
	"strconv"

	"github.com/issueye/common/pkg/utils"
)

type Client struct {
	Base
	ClientBase
}

// ClientBase
// 客户端信息
type ClientBase struct {
	Name         string `label:"名称" gorm:"column:name;type:nvarchar(200);comment:名称;" json:"name"`                                       // 名称
	Title        string `label:"标题" gorm:"column:title;type:nvarchar(200);comment:标题;" json:"title"`                                     // 标题
	IP           string `label:"IP" gorm:"column:ip;type:nvarchar(200);comment:IP;" json:"ip"`                                           // IP
	Host         string `label:"主机" gorm:"column:host;type:nvarchar(200);comment:主机;" json:"host"`                                       // 主机
	OS           uint   `label:"操作系统" gorm:"column:os;type:int;comment:操作系统;" json:"os"`                                                 // 操作系统 0 windows 1 linux
	State        uint   `label:"状态" gorm:"column:state;type:int;comment:状态;" json:"state"`                                               // 状态 0 停用 1 启用
	Version      string `label:"版本" gorm:"column:version;type:nvarchar(200);comment:版本;" json:"version"`                                 // 版本
	GitHash      string `label:"GitHash" gorm:"column:git_hash;type:nvarchar(200);comment:GitHash;" json:"git_hash"`                     // GitHash
	GoVersion    string `label:"GoVersion" gorm:"column:go_version;type:nvarchar(200);comment:GoVersion;" json:"go_version"`             // GoVersion
	PrintVersion string `label:"PrintVersion" gorm:"column:print_version;type:nvarchar(200);comment:PrintVersion;" json:"print_version"` // PrintVersion
}

func (mod *ClientBase) Copy(data *ClientBase) {
	mod.Name = data.Name
	mod.Title = data.Title
	mod.IP = data.IP
	mod.Host = data.Host
	mod.OS = data.OS
	mod.State = data.State
}

// TableName
// 表名称
func (Client) TableName() string {
	return "client_info"
}

func (Client) New() *Client {
	return &Client{
		Base: Base{
			ID: strconv.FormatInt(utils.GenID(), 10),
		},
	}
}
