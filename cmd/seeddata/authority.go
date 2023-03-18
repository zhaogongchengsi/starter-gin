package seeddata

import (
	"github.com/server-gin/module"
	"gorm.io/gorm"
)

func CreateAuthorityTable(db *gorm.DB) error {
	return db.AutoMigrate(module.Authority{})
}

var Authoritys = []module.Authority{
	{
		AuthorityName: "超级管理员",
		AuthorityId:   1,
		ParentId:      0,
		RouterRecords: Routers,
	},
	{
		AuthorityName: "二级管理员",
		AuthorityId:   2,
		ParentId:      1,
	},
}

func CrateAuthouitySeedData(db *gorm.DB) error {
	return db.Model(module.Authority{}).Create(Authoritys).Error
}
