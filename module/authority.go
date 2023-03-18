package module

type Authority struct {
	BaseMode
	AuthorityId   int            `json:"authorityId" gorm:"not null;unique;primary_key;comment:角色ID;size:90"`
	AuthorityName string         `json:"authorityName" gorm:"comment:角色名"`
	ParentId      int            `json:"parentId" gorm:"comment:父角色ID"`
	Children      []Authority    `json:"children" gorm:"-"`
	RouterRecords []RouterRecord `json:"routerRecords" gorm:"many2many:authority_routers;"`
}
