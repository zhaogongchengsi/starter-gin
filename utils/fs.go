package utils

import (
	"io"
	"mime/multipart"
	"os"
	"path"
	"path/filepath"
)

// dir 需要保存的目录
func SaveFileHeader(fh *multipart.FileHeader, dir string) (string, error) {

	src, err := fh.Open()
	if err != nil {
		return "", err
	}
	defer src.Close()

	filename := fh.Filename
	fileext := filepath.Ext(filename)

	filename = MD5([]byte(filename)) + fileext

	target := path.Join(dir, filename)

	if err = os.MkdirAll(filepath.Dir(dir), os.ModePerm); err != nil {
		return "", err
	}

	out, err := os.Create(target)
	if err != nil {
		return "", err
	}
	defer out.Close()

	_, err = io.Copy(out, src)
	if err != nil {
		return "", err
	}

	return filename, nil

}
