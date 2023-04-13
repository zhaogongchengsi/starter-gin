package actions

import (
	"github.com/zhaogongchengsi/starter-gin/cmd/tools/options"
	"github.com/zhaogongchengsi/starter-gin/cmd/tools/seeddata"
	"github.com/zhaogongchengsi/starter-gin/module"
	"github.com/zhaogongchengsi/starter-gin/utils"
	"os"
)

func InitAction() {

	err := initApp()
	if err != nil {
		panic(err)
	}

	utils.Success("Application initialized successfully")

	os.Exit(0)
}

func initApp() error {
	db, err := ConnDb(options.Option.Config)

	if err != nil {
		return err
	}

	err = db.SetupJoinTable(&module.User{}, "Authorities", &module.UserAuthority{})
	err = db.AutoMigrate(module.User{}, module.File{}, &module.Languages{}, &module.Language{}, module.RouterRecord{})
	if err != nil {
		return err
	}

	db.Create(seeddata.Users)
	return nil
}
