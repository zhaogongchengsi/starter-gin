package cmd

import (
	"fmt"
	"os"
	"strings"

	"github.com/server-gin/modules/system"
)

func autoMigAction(ms string) {
	modules := strings.Split(ms, ",")
	err := AutoMigrateModule(modules)
	if err != nil {
		panic(err)
	}
	os.Exit(0)
}

var moduleMap = map[string]any{
	"user":   system.User{},
	"file":   system.File{},
	"router": system.RouterRecord{},
}

func AutoMigrateModule(ms []string) error {

	db, err := ConnDb(*c, *t, "config")
	if err != nil {
		return err
	}
	for _, v := range ms {
		md, ok := moduleMap[strings.TrimSpace(v)]
		if !ok {
			fmt.Printf("%s model does not exist", v)
			continue
		}

		err := db.AutoMigrate(md)
		if err != nil {
			fmt.Printf("autoMigrate Error: %v", err.Error())
		}
	}

	return nil
}
