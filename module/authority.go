package module

import (
	"errors"
	"fmt"
	"github.com/zhaogongchengsi/starter-gin/common"
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
	RouterRecords []RouterRecord `json:"routerRecords" gorm:"many2many:authority_routers;"`
}

func (*Authority) TableName() string {
	return "authority"
}

// AuthorityIdKey id
func (*Authority) AuthorityIdKey() string {
	return "id"
}

// RouterRecordRelevancyKey "RouterRecords"
func (*Authority) RouterRecordRelevancyKey() string {
	return "RouterRecords"
}

func NewAuthority(authorityId int) *Authority {
	return &Authority{Id: authorityId}
}

func NewFullAuthority(authorityId, authorityPid int, name string) *Authority {
	return &Authority{Id: authorityId, ParentId: authorityPid, AuthorityName: name}
}

func (a *Authority) GetAuthorities(page common.Page, db *gorm.DB) (authorities []Authority, err error) {
	err = db.Scopes(common.Paginate(page)).Find(&authorities).Error
	return authorities, err
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

func (a *Authority) DeleteAuthority(ids []int, db *gorm.DB) error {

	return db.Transaction(func(tx *gorm.DB) error {
		var auths []Authority
		err := db.Model(a).Where(ids).Preload(a.RouterRecordRelevancyKey()).Find(&auths).Error
		if err != nil {
			return err
		}
		if len(auths) == 0 {
			return nil
		}

		err = db.Model(a).Delete(&auths).Error
		if err != nil {
			return err
		}

		ua := new(UserAuthority)

		err = ua.DeleteUserAuthsByAuthIds(ids, *tx)
		if err != nil {
			return err
		}
		//var rs []RouterRecord
		//for _, auth := range auths {
		//	rs = append(rs, auth.RouterRecords...)
		//}
		//
		//for _, r := range rs {
		//	fmt.Printf("id: %v, name: %s\n", r.ID, r.Name)
		//}
		//
		//err = db.Model(a).Association(a.RouterRecordRelevancyKey()).Delete(&rs)
		//if err != nil {
		//	return err
		//}
		return nil
	})
}
