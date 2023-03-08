package cmd

import (
	"fmt"
	"os"
	"strings"

	"github.com/server-gin/core"
	"github.com/server-gin/global"
	"github.com/server-gin/modules/system"
	"gorm.io/gorm"
)

func seedAction(file string) {
	ags := strings.Split(file, "-")

	var p = "./"
	var t = "yaml"
	var n = "config"
	if len(ags[0]) != 0 {
		p = ags[0]
	}
	if len(ags[1]) != 0 {
		t = ags[1]
	}
	if len(ags[2]) != 0 {
		n = ags[2]
	}

	err := seed(p, t, n)

	if err != nil {
		panic(err)
	}
	os.Exit(0)
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
