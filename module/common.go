package module

import (
	"errors"
	"gorm.io/gorm"
	"time"
)

type BaseMode struct {
	ID        uint           `gorm:"primarykey" json:"id"` // 主键ID
	CreatedAt time.Time      `json:"createAt"`             // 创建时间
	UpdatedAt time.Time      `json:"updateAt"`             // 更新时间
	DeletedAt gorm.DeletedAt `gorm:"index" json:"-"`       // 删除时间
}

var (
	ErrUserNotExist = errors.New("err: user does not exist")
	// ErrUserAuthExists 关系已存在
	ErrUserAuthExists   = errors.New("err: relationship already exists")
	ErrUserAuthNotExist = errors.New("err: relationship does not exist")
	ErrAuthExist        = errors.New("err: authority already exists")
	ErrAuthNotExist     = errors.New("err: authority does not exist")
)
