package actions

import (
	"fmt"
	"github.com/zhaogongchengsi/starter-gin/cmd/tools/options"
	"github.com/zhaogongchengsi/starter-gin/utils"
	"os"
	"strings"
)

func AutoMigAction(ms string) {
	modules := strings.Split(ms, ",")
	err := AutoMigrateModule(modules)
	if err != nil {
		panic(err)
	}
	utils.Success("Migration successful!")
	os.Exit(0)
}

func AutoMigrateModule(ms []string) error {

	db, err := ConnDb(options.Option.Config)
	if err != nil {
		return err
	}

	if name := strings.TrimSpace(ms[0]); name == "all" {
		for _, fun := range moduleMap {
			err := fun(db)
			if err != nil {
				return fmt.Errorf("autoMigrate Error: %v", err.Error())
			}
		}
		return nil
	}

	for _, v := range ms {
		name := strings.TrimSpace(v)

		defunct, ok := moduleMap[name]
		if !ok {
			utils.Warning("%s model does not exist", v)
			continue
		}

		err := defunct(db)
		if err != nil {
			utils.Warning("autoMigrate Error: %v", err.Error())
		}
	}

	return nil
}
