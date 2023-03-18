package cmd

import (
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
		for _, seedfun := range moduleSeedMap {
			if err := seedfun(db); err != nil {
				return err
			}
		}
		return nil
	}

	for _, v := range ms {
		name := strings.TrimSpace(v)

		mdc, ok := moduleSeedMap[name]
		if !ok {
			fmt.Printf("%s model does not exist", v)
			continue
		}

		err := mdc(db)
		if err != nil {
			return err
		}
	}

	return nil
}
