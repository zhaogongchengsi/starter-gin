package cmd

import (
	"errors"
	"fmt"
	"os"
	"strings"
)

func seedAction(file string) {
	ags := strings.Split(file, ",")

	err := seed(ags)

	if err != nil {
		panic(err)
	}
	os.Exit(0)
}

func seed(ms []string) error {

	db, err := ConnDb(*c, *t, *n)

	if err != nil {
		return fmt.Errorf("seed Error: Database connection failed, please check %s and try again", err)
	}

	if name := strings.TrimSpace(ms[0]); name == "all" {
		for name, data := range moduleSeedMap {
			module, ok := moduleMap[name]
			if !ok {
				fmt.Printf("%s 不存在", name)
				return errors.New(name + "不存在")
			}

			if err := db.Model(module).Create(data).Error; err != nil {
				return err
			}
		}
		return nil
	}

	for _, v := range ms {
		name := strings.TrimSpace(v)

		md, ok := moduleSeedMap[name]
		if !ok {
			fmt.Printf("%s model does not exist", v)
			continue
		}

		err := db.Create(md).Error
		if err != nil {
			return err
		}
	}

	return nil
}
