package options

type Options struct {
	Config  string       `short:"c" long:"config" description:"the configuration file" default:"./config.yaml"`
	Seed    func(string) `short:"s" long:"seed" description:"filePath-fileType-fileName"`
	Ssl     func(string) `short:"g" long:"gsc" description:"Generate ssl certificate"`
	AutoMig func(string) `short:"a" long:"auto" description:"Initialize model"`
	Init    func()       `short:"i" long:"init" description:"Initialize Apply initialize model and insert seed data"`
}

var Option *Options = &Options{
	Config: "./config.yaml",
}
