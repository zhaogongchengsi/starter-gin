package cmd

import (
	"flag"
	"os"

	"github.com/jessevdk/go-flags"
	"github.com/zhaogongchengsi/starter-gin/global"
)

var c = flag.String("c", "./", "Directory where configuration files are stored")
var t = flag.String("t", "yaml", "Type of configuration file")
var n = flag.String("n", "config", "Name of the configuration file")

type Options struct {
	ConfigDir  string       `short:"c" long:"config" description:"Directory where configuration files are stored" default:"./"`
	ConfigType string       `short:"t" long:"configType" description:"Type of configuration file" default:"yaml"`
	ConfigName string       `short:"n" long:"configName" description:"Name of the configuration file" default:"config"`
	Mode       string       `short:"m" long:"mode" description:"gin.mode:[debug, release, test] " default:"debug"`
	Seed       func(string) `short:"s" long:"seed" description:"filePath-fileType-fileName"`
	Ssl        func(string) `short:"g" long:"gsc" description:"Generate ssl certificate"`
	AutoMig    func(string) `short:"a" long:"auto" description:"Initialize model"`
}

func Parse() error {

	var opt Options
	opt.Seed = seedAction
	opt.Ssl = sslAction
	opt.AutoMig = autoMigAction

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
