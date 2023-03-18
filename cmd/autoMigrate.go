package cmd

import (
	"fmt"
	"os"
	"strings"
)

func autoMigAction(ms string) {
	modules := strings.Split(ms, ",")
	err := AutoMigrateModule(modules)
	if err != nil {
		panic(err)
	}
	os.Exit(0)
}

func AutoMigrateModule(ms []string) error {

	db, err := ConnDb(*c, *t, *n)
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
			fmt.Printf("%s model does not exist", v)
			continue
		}

		err := defunct(db)
		if err != nil {
			fmt.Printf("autoMigrate Error: %v", err.Error())
		}
	}

	return nil
}
