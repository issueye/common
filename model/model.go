package model

import "time"

type Page[T any] struct {
	PageNum   int64 `json:"pageNum" form:"pageNum"`   // 页数
	PageSize  int64 `json:"pageSize" form:"pageSize"` // 页码
	Total     int64 `json:"total"`                    // 总数  由服务器返回回去
	Condition T     `json:"condition"`                // 条件
}

func NewPage[T any](condition T) *Page[T] {
	return &Page[T]{Condition: condition}
}

type Base struct {
	ID        string    `gorm:"column:id;size:100;comment:编码;primaryKey;autoIncrement:false;" json:"id"` // 编码
	CreatedAt time.Time `gorm:"column:created_at;type:datetime;comment:创建时间;" json:"createdAt"`          // 创建时间
	UpdatedAt time.Time `gorm:"column:updated_at;type:datetime;comment:修改时间;" json:"updatedAt"`          // 修改时间
}
