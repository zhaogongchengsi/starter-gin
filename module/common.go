package module

import (
	"gorm.io/gorm"
	"time"
)

type BaseMode struct {
	ID        uint           `gorm:"primarykey" json:"id"` // 主键ID
	CreatedAt time.Time      `json:"createAt"`             // 创建时间
	UpdatedAt time.Time      `json:"updateAt"`             // 更新时间
	DeletedAt gorm.DeletedAt `gorm:"index" json:"-"`       // 删除时间
}
