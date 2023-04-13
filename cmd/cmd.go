package cmd

import (
	"os"

	"github.com/jessevdk/go-flags"
	"github.com/zhaogongchengsi/starter-gin/global"
)

type Options struct {
	ConfigDir  string       `short:"c" long:"config" description:"Directory where configuration files are stored" default:"./"`
	ConfigType string       `short:"t" long:"configType" description:"Type of configuration file" default:"yaml"`
	ConfigName string       `short:"n" long:"configName" description:"Name of the configuration file" default:"config"`
	Mode       string       `short:"m" long:"mode" description:"gin.mode:[debug, release, test] " default:"debug"`
	Seed       func(string) `short:"s" long:"seed" description:"filePath-fileType-fileName"`
	Ssl        func(string) `short:"g" long:"gsc" description:"Generate ssl certificate"`
	AutoMig    func(string) `short:"a" long:"auto" description:"Initialize model"`
	Init       func()       `short:"i" long:"init" description:"Initialize Apply initialize model and insert seed data"`
}

var Opt Options

func Parse() error {

	//var opt Options
	Opt.Seed = seedAction
	Opt.Ssl = sslAction
	Opt.AutoMig = autoMigAction
	Opt.Init = InitAction

	_, err := flags.Parse(&Opt)

	if isH := flags.WroteHelp(err); isH {
		os.Exit(0)
	}

	if err != nil {
		return err
	}

	global.ConfigDirPath = Opt.ConfigDir
	global.ConfigName = Opt.ConfigName
	if ok := include([]string{"yaml", "json"}, Opt.ConfigType); !ok {
		global.ConfigType = "yaml"
	} else {
		global.ConfigType = Opt.ConfigType
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
