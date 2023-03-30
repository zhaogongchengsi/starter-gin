package module

import (
	"errors"
	"fmt"
	"gorm.io/gorm"
)

type UserAuthority struct {
	UserID      int `gorm:"primaryKey"`
	AuthorityId int `gorm:"primaryKey"`
	DeletedAt   gorm.DeletedAt
}

func (UserAuthority) UserIdKey() string {
	return "user_id"
}
func (UserAuthority) UserAuthorityIdKey() string {
	return "authority_id"
}

func (UserAuthority) TableName() string {
	return "user_and_authorities"
}

func NewUserAuthority(userId int, authorityId int) *UserAuthority {
	return &UserAuthority{UserID: userId, AuthorityId: authorityId}
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

func (ua UserAuthority) DeleteUserAuth(db *gorm.DB) error {
	err := db.Model(ua).Where(fmt.Sprintf("%s = ? AND %s = ?", ua.UserIdKey(), ua.UserAuthorityIdKey()), ua.UserID, ua.AuthorityId).Unscoped().Delete(&ua).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return ErrUserAuthNotExist
		}
		return err
	}
	return nil
}

func (ua UserAuthority) DeleteUserAuthsByAuthIds(ids []int, db gorm.DB) error {
	var auths []UserAuthority
	err := db.Table(ua.TableName()).Where(ua.UserAuthorityIdKey()+" in ?", ids).Find(&auths).Error
	if err != nil {
		return err
	}
	return db.Table(ua.TableName()).Delete(&auths).Error
}
