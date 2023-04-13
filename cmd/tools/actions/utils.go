package actions

import (
	"fmt"
	seedData "github.com/zhaogongchengsi/starter-gin/cmd/tools/seeddata"
	"github.com/zhaogongchengsi/starter-gin/config"

	"github.com/zhaogongchengsi/starter-gin/core"
	"gorm.io/gorm"
)

type CreateFunc = func(db *gorm.DB) error

var moduleMap = map[string]CreateFunc{
	"user":      seedData.CreateUserTable,
	"file":      seedData.CrateFileTable,
	"router":    seedData.CreateRouter,
	"languages": seedData.CreateLanguages,
	"language":  seedData.CreateLanguages,
	//"authority": seeddata.CreateAuthorityTable,
}

var moduleSeedMap = map[string]CreateFunc{
	"user": seedData.CrateUserSeedData,
	//"router":    seeddata.CrateRouterSeedData,
	"languages": seedData.CrateLanguagesSeedData,
	//"auths":     seeddata.CrateAuthouitySeedData,
}

func ConnDb(file string) (*gorm.DB, error) {

	conf, err := config.LoadServerConfig(file)

	if err != nil {
		return &gorm.DB{}, fmt.Errorf("seed Error: The specified parameters are wrong, and the database configuration cannot be obtained. %s %v", file, err)
	}

	db, err := core.ConnectDataBaseServer(conf)

	if err != nil {
		return &gorm.DB{}, fmt.Errorf("seed Error: Database connection failed, please check %s and try again", err)
	}

	return db, nil
}

func include[T string | int | int64 | int32](arr []T, target T) bool {
	count := len(arr)
	exist := false
	for i := 0; i < count; i++ {
		item := arr[i]
		if item == target {
			exist = true
		}
	}
	return exist
}
