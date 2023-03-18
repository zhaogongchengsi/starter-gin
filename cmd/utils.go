package cmd

import (
	"fmt"
	"github.com/server-gin/cmd/seeddata"

	"github.com/server-gin/core"
	"github.com/server-gin/global"
	"gorm.io/gorm"
)

type CreateFunc = func(db *gorm.DB) error

var moduleMap = map[string]CreateFunc{
	"user":      seeddata.CreateUserTable,
	"file":      seeddata.CrateFileTable,
	"router":    seeddata.CreateRouter,
	"languages": seeddata.CreateLanguages,
	"language":  seeddata.CreateLanguages,
	"authority": seeddata.CreateAuthorityTable,
}

var moduleSeedMap = map[string]CreateFunc{
	"user":      seeddata.CrateUserSeedData,
	"router":    seeddata.CrateRouterSeedData,
	"languages": seeddata.CrateLanguagesSeedData,
	"auths":     seeddata.CrateAuthouitySeedData,
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
