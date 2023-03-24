package module

import (
	"errors"
	"fmt"
	"gorm.io/gorm"
	"time"
)

var (
	ErrAuthExist    = errors.New("err: already exists")
	ErrAuthNotExist = errors.New("err: does not exist")
)

type Authority struct {
	CreatedAt     time.Time      `json:"createAt"`       // 创建时间
	UpdatedAt     time.Time      `json:"updateAt"`       // 更新时间
	DeletedAt     gorm.DeletedAt `gorm:"index" json:"-"` // 删除时间
	AuthorityId   int            `json:"authorityId" gorm:"not null;unique;primary_key;comment:角色ID;size:90"`
	AuthorityName string         `json:"authorityName" gorm:"comment:角色名"`
	ParentId      int            `json:"parentId" gorm:"comment:父角色ID"`
	Users         []User         `json:"users" gorm:"many2many:user_authoritys;"`
	Children      []Authority    `json:"children" gorm:"-"`
	RouterRecords []RouterRecord `json:"routerRecords" gorm:"many2many:authority_routers;"`
}

func (*Authority) AuthorityIdKey() string {
	return "authority_id"
}

func NewAuthority(authorityId int) *Authority {
	return &Authority{AuthorityId: authorityId}
}

func NewFullAuthority(authorityId, authorityPid int, name string) *Authority {
	return &Authority{AuthorityId: authorityId, ParentId: authorityPid, AuthorityName: name}
}

func (a *Authority) GetUserAuths(uid string, db *gorm.DB) (User, error) {
	var user User
	err := db.Model(&User{}).Preload(user.AuthRelevancyKey()).Where("uuid = ?", uid).First(&user).Error
	return user, err
}

func (a *Authority) CreateAuth(db *gorm.DB) error {
	fmt.Println(a.AuthorityId, a.AuthorityName, a.ParentId)
	err := db.Create(&a).Error
	if err != nil {
		if errors.Is(err, gorm.ErrDuplicatedKey) {
			return ErrAuthExist
		}
		return err
	}
	return nil
}

func (a *Authority) FindAuthority(db *gorm.DB) error {
	var au Authority
	err := db.Model(a).Where(fmt.Sprintf("%s = ?", a.AuthorityIdKey()), a.AuthorityId).First(&au).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return ErrAuthNotExist
		}
		return err
	}
	return nil
}
