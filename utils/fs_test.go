package utils_test

import (
	"os"
	"testing"

	"github.com/server-gin/utils"
)

func TestFsCreateFile(t *testing.T) {
	root, err := os.Getwd()
	if err != nil {
		t.Error(err)
	}

	path := root + "/test.txt"
	context := "test"

	err = utils.CreateFile(path, context)

	if err != nil {
		t.Error(err)
	}

	data, err := utils.ReadFile(path)

	if err != nil {
		t.Error(err)
	}

	if data != context {
		t.Errorf("file content is incorrect : %s", data)
	}

	os.Remove(path)

}
