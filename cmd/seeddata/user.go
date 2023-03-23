package seeddata

import (
	uuid "github.com/satori/go.uuid"
	"github.com/zhaogongchengsi/starter-gin/module"
	"github.com/zhaogongchengsi/starter-gin/utils"
	"gorm.io/gorm"
)

func CreateUserTable(db *gorm.DB) error {
	return db.AutoMigrate(module.User{}, module.Authority{})
}

var Users = []module.User{
	{
		Phone:      "12312312312",
		Password:   module.CreatePassworld(utils.MD5([]byte("123456"))),
		UUID:       uuid.NewV4(),
		UserName:   "admin",
		NickName:   "超级管理员",
		Authoritys: Authoritys,
	},
}

func CrateUserSeedData(db *gorm.DB) error {
	return db.Model(module.User{}).Create(Users).Error
}
