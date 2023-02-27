package utils_test

import (
	"os"
	"path"
	"testing"

	"github.com/server-gin/utils"
)

func TestNewCertCnfig(t *testing.T) {
	c := utils.NewCertConifg("localhost")
	_, err := c.CreateCertificate()

	if err != nil {
		t.Error(err)
	}

	_, err = c.CreateKey()

	if err != nil {
		t.Error(err)
	}
}

func TestGrCertConfig(t *testing.T) {
	root, err := os.Getwd()
	if err != nil {
		t.Error(err)
	}
	c := utils.NewCertConifg("localhost")
	err = c.Generate(root, "abc", "abckey")
	if err != nil {
		t.Error(err)
	}

	certPath := path.Join(root, "abc.pem")
	keyPath := path.Join(root, "abckey.pem")
	os.Remove(certPath)
	os.Remove(keyPath)
}
