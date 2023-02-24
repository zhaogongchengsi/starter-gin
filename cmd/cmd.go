package cmd

import (
	"fmt"
	"os"

	"github.com/jessevdk/go-flags"
	"github.com/server-gin/global"
)

type Options struct {
	ConfigDir  string       `short:"c" long:"config" description:"Directory where configuration files are stored" default:"configs"`
	ConfigType string       `short:"t" long:"configType" description:"Type of configuration file" default:"yaml"`
	Init       func(string) `short:"i" long:"init" description:"Initialize the database seed data parameter to database url"`
}

func Parse() error {
	var opt Options

	opt.Init = func(dns string) {
		fmt.Println("执行初始化")
	}

	_, err := flags.Parse(&opt)

	if isH := flags.WroteHelp(err); isH {
		os.Exit(0)
	}

	if err != nil {
		fmt.Println(err)
		return err
	}

	global.ConfigDirPath = opt.ConfigDir
	global.ConfigType = opt.ConfigType

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
