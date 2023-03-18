package seeddata

import (
	uuid "github.com/satori/go.uuid"
	"github.com/server-gin/modules/system"
	"gorm.io/gorm"
)

func CreateUserTable(db *gorm.DB) error {
	err := db.AutoMigrate(system.User{})
	if err != nil {
		return err
	}
	return CreateAuthorityTable(db)
}

var Users = []system.User{
	{
		Phone:    "12312312312",
		Password: system.CreatePassworld("123456"),
		UUID:     uuid.NewV4(),
		UserName: "admin",
		NickName: "超级管理员",
	},
}

func CrateUserSeedData(db *gorm.DB) error {
	return db.Model(system.User{}).Create(Users).Error
}
