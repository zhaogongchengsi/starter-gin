package module

import "gorm.io/gorm"

type UserAuthority struct {
	UserId      int `gorm:"column:user_id"`
	AuthorityId int `gorm:"column:authority_authority_id"`
}

func (UserAuthority) TableName() string {
	return "user_authoritys"
}

func NewUserAuthority(userId int, authorityId int) *UserAuthority {
	return &UserAuthority{UserId: userId, AuthorityId: authorityId}
}

func (ua UserAuthority) CreateUserAuth(db *gorm.DB) error {
	return db.Model(ua).Create(&ua).Error
}
