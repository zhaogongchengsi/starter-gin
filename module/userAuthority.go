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

func (ua UserAuthority) FindUserAuthorityByAuthIds(ids []int, db *gorm.DB) ([]UserAuthority, error) {
	var auths []UserAuthority
	err := db.Table(ua.TableName()).Where(ua.UserAuthorityIdKey()+" in ?", ids).Find(&auths).Error
	if err != nil {
		return auths, err
	}
	return auths, err
}

func (ua UserAuthority) FirstUserAuthorityByAuthId(id int, db *gorm.DB) (a UserAuthority, e error) {
	e = db.Table(ua.TableName()).Where(ua.UserAuthorityIdKey()+" = ?", id).First(&a).Error
	if e != nil {
		return a, e
	}
	return a, e
}

// WhetherByAuthId 通过权限id查询是否分配给某个角色  为 false 没有分配   true 为分配
func (ua UserAuthority) WhetherByAuthId(id int, db *gorm.DB) bool {
	return errors.Is(db.Model(ua).Where(ua.UserAuthorityIdKey()+" = ?", id).Error, gorm.ErrRecordNotFound)
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

func (ua UserAuthority) DeleteUserAuthsByAuthId(id int, db *gorm.DB) error {
	return db.Table(ua.TableName()).Delete(&[]UserAuthority{}, ua.UserAuthorityIdKey()+" = ?", id).Error
}
