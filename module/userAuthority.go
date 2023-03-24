package module

import (
	"errors"
	"gorm.io/gorm"
)

type UserAuthority struct {
	UserId      int `gorm:"column:user_id"`
	AuthorityId int `gorm:"column:authority_authority_id"`
}

var (
	// ErrUserAuthExists 关系已存在
	ErrUserAuthExists = errors.New("err: relationship already exists")
)

func (UserAuthority) TableName() string {
	return "user_authoritys"
}

func NewUserAuthority(userId int, authorityId int) *UserAuthority {
	return &UserAuthority{UserId: userId, AuthorityId: authorityId}
}

func (ua UserAuthority) CreateUserAuth(db *gorm.DB) error {
	err := db.Model(ua).Create(&ua).Error
	if err != nil {
		if errors.Is(err, gorm.ErrDuplicatedKey) {
			return ErrUserAuthExists
		}
		return err
	}
	return nil
}
