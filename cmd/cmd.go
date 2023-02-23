package cmd

import (
	"flag"
)

func Parse(configDir *string, configType *string) {
	defaultType := "yaml"
	flag.StringVar(configDir, "config", "./configs", "Store configuration directory")
	flag.StringVar(configType, "configType", defaultType, "Type of configuration file")

	if is := include([]string{"yaml", "json"}, *configType); !is {
		configType = &defaultType
	}

	flag.Parse()
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
