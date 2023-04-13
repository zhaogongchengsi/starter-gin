package main

import (
	"github.com/jessevdk/go-flags"
	"github.com/zhaogongchengsi/starter-gin/cmd/tools/actions"
	"github.com/zhaogongchengsi/starter-gin/cmd/tools/options"
	"os"
)

func main() {
	//var opt Options
	options.Option.Seed = actions.SeedAction
	options.Option.Ssl = actions.SslAction
	options.Option.AutoMig = actions.AutoMigAction
	options.Option.Init = actions.InitAction

	_, err := flags.Parse(options.Option)

	if isH := flags.WroteHelp(err); isH {
		os.Exit(0)
	}

	if err != nil {
		panic(err)
	}

}
