package config_test

import (
	"fmt"
	"os"
	"path"
	"strings"
	"testing"

	"github.com/server-gin/config"
)

func formatPath(path string) string {
	return strings.ReplaceAll(path, "\\", "/")
}

func TestReadConfig(t *testing.T) {

	root, err := os.Getwd()

	if err != nil {
		t.Error(err)
	}

	p := formatPath(path.Join(root + "s"))

	c := config.NewConfig(p, "yaml", "server", "database")

	err = c.ReadConfigs()

	if err != nil {
		t.Error(err)
	}

}

func TestReadValue(t *testing.T) {

	root, err := os.Getwd()

	if err != nil {
		t.Error(err)
	}

	p := formatPath(path.Join(root + "s"))
	c := config.NewConfig(p, "yaml", "server")
	err = c.ReadConfigs()

	if err != nil {
		t.Error(err)
	}

	ser := config.Server{}

	err = c.ReadValue(&ser, "Server")

	if err != nil {
		t.Error(err)
	}

	fmt.Println(ser)

	if ser.Port == 0 {
		t.Errorf("Read as null")
	}

}
