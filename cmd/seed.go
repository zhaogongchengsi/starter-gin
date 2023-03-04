package cmd

import (
	"fmt"

	"github.com/server-gin/core"
	"github.com/server-gin/global"
	"github.com/server-gin/modules/system"
)

func seed(file, typ, name string) error {

	conf, err := global.ReadConfig(file, typ, name)

	if err != nil {
		return fmt.Errorf("seed Error: The specified parameters are wrong, and the database configuration cannot be obtained. %s %s %v", file, typ, err)
	}

	db, err := core.ConnectDataBaseServer(conf)

	if err != nil {
		return fmt.Errorf("seed Error: Database connection failed, please check %s and try again", err)
	}

	var user system.User

	err = db.AutoMigrate(user)

	if err != nil {
		return fmt.Errorf("seed Error: database initialization failed : %v", err)
	}

	newUser := system.NewUser("admin", "12345", "18312391231", "管理员", "zzh123123123@qq.com")

	db.Create(newUser)

	return nil
}
