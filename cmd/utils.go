package cmd

import (
	"fmt"

	uuid "github.com/satori/go.uuid"
	"github.com/server-gin/core"
	"github.com/server-gin/global"
	"github.com/server-gin/modules/system"
	"gorm.io/gorm"
)

var moduleMap = map[string]any{
	"user":   system.User{},
	"file":   system.File{},
	"router": system.RouterRecord{},
}

var moduleSeedMap = map[string]any{
	"user": []system.User{
		{
			Phone:    "12312312312",
			Password: system.CreatePassworld("123456"),
			UUID:     uuid.NewV4(),
			UserName: "admin",
			NickName: "超级管理员",
		},
	},
}

func ConnDb(file, typ, name string) (*gorm.DB, error) {

	conf, err := global.ReadConfig(file, typ, name)

	if err != nil {
		return &gorm.DB{}, fmt.Errorf("seed Error: The specified parameters are wrong, and the database configuration cannot be obtained. %s %s %v", file, typ, err)
	}

	db, err := core.ConnectDataBaseServer(conf)

	if err != nil {
		return &gorm.DB{}, fmt.Errorf("seed Error: Database connection failed, please check %s and try again", err)
	}

	return db, nil
}
