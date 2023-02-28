package cmd

import (
	"os"

	"github.com/jessevdk/go-flags"
	"github.com/server-gin/global"
)

type Options struct {
	ConfigDir  string       `short:"c" long:"config" description:"Directory where configuration files are stored" default:"configs"`
	ConfigType string       `short:"t" long:"configType" description:"Type of configuration file" default:"yaml"`
	Seed       func(string) `short:"s" long:"seed" description:"Initialize the database seed data parameter to database url"`
	Ssl        func(string) `short:"g" long:"gsc" description:"Generate ssl certificate"`
}

func Parse() error {
	var opt Options
	opt.Seed = func(s string) {
		err := seed(s)
		if err != nil {
			panic(err)
		}
		os.Exit(0)
	}

	opt.Ssl = func(s string) {
		err := generateSsl(s)
		if err != nil {
			panic(err)
		}
		os.Exit(0)
	}

	_, err := flags.Parse(&opt)

	if isH := flags.WroteHelp(err); isH {
		os.Exit(0)
	}

	if err != nil {
		return err
	}

	global.ConfigDirPath = opt.ConfigDir
	if ok := include([]string{"yaml", "json"}, opt.ConfigType); !ok {
		global.ConfigType = "yaml"
	} else {
		global.ConfigType = opt.ConfigType
	}

	return nil
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

// root:''@tcp(localhost:3306)/starter_gin?charset=utf8mb4&parseTime=True&loc=Local
