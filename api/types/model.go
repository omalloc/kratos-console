package types

import (
	"gorm.io/gorm"
	"time"
)

type QueryKVFilter struct {
	Key   string `json:"key"`
	Value string `json:"value"`
}

type DBModel struct {
	CreatedAt time.Time      `json:"created_at" gorm:"column:created_at;comment:创建时间"`
	UpdatedAt time.Time      `json:"updated_at" gorm:"column:updated_at;comment:更新时间"`
	DeletedAt gorm.DeletedAt `json:"-" gorm:"index;column:deleted_at;comment:删除时间"`
}
