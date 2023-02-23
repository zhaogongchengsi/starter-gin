package cmd

import (
	"flag"
)

func ParseServerOptions(configDir *string, configType *string) {
	defaultType := "yaml"
	flag.StringVar(configDir, "config", "./configs", "Store configuration directory")
	flag.StringVar(configType, "configType", defaultType, "Type of configuration file")

	if is := include([]string{"yaml", "json"}, *configType); !is {
		configType = &defaultType
	}
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

func ParseDevOptions(init *bool) {
	flag.BoolVar(init, "init", false, "Whether to initialize application data")
}
