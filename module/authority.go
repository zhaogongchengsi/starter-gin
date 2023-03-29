package module

import (
	"errors"
	"fmt"
	"gorm.io/gorm"
	"time"
)

type Authority struct {
	CreatedAt     time.Time      `json:"createAt"`       // 创建时间
	UpdatedAt     time.Time      `json:"updateAt"`       // 更新时间
	DeletedAt     gorm.DeletedAt `gorm:"index" json:"-"` // 删除时间
	Id            int            `json:"id" gorm:"not null;unique;primaryKey;comment:角色ID;size:90"`
	AuthorityName string         `json:"authorityName" gorm:"comment:角色名"`
	ParentId      int            `json:"pid" gorm:"comment:父角色ID"`
	//Users         []User         `json:"users" gorm:"many2many:user_and_authorities;"`
	Children      []Authority    `json:"children" gorm:"-"`
	RouterRecords []RouterRecord `json:"routerRecords" gorm:"many2many:authority_routers;"`
}

func (*Authority) TableName() string {
	return "authority"
}

func (*Authority) AuthorityIdKey() string {
	return "authority_id"
}

func NewAuthority(authorityId int) *Authority {
	return &Authority{Id: authorityId}
}

func NewFullAuthority(authorityId, authorityPid int, name string) *Authority {
	return &Authority{Id: authorityId, ParentId: authorityPid, AuthorityName: name}
}

func (a *Authority) GetUserAuths(uid string, db *gorm.DB) (User, error) {
	var user User
	err := db.Model(&User{}).Preload(user.AuthRelevancyKey()).Where("uuid = ?", uid).First(&user).Error
	return user, err
}

func (a *Authority) CreateAuth(db *gorm.DB) error {
	fmt.Println(a.Id, a.AuthorityName, a.ParentId)
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
	err := db.Model(a).Where(fmt.Sprintf("%s = ?", a.AuthorityIdKey()), a.Id).First(&au).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return ErrAuthNotExist
		}
		return err
	}
	return nil
}
