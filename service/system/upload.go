package system

import (
	"io"
	"mime/multipart"
	"os"
	"path"
	"path/filepath"

	"github.com/server-gin/global"
	"github.com/server-gin/modules/system"
	"github.com/server-gin/utils"
)

type UpdateFileInfo struct {
	FileName   string                `json:"fileName"`
	FileHeader *multipart.FileHeader `json:"-"`
}

func (u *UpdateFileInfo) SaveFile() (system.File, error) {

	src, err := u.FileHeader.Open()
	if err != nil {
		return system.File{}, err
	}
	defer src.Close()

	dir := path.Join(global.AppConfig.Server.UploadDir, utils.MD5([]byte(u.FileName))+filepath.Ext(u.FileName))

	if err = os.MkdirAll(filepath.Dir(dir), 0750); err != nil {
		return system.File{}, err
	}

	out, err := os.Create(dir)
	if err != nil {
		return system.File{}, err
	}
	defer out.Close()

	_, err = io.Copy(out, src)
	if err != nil {
		return system.File{}, err
	}

	return system.File{}, nil
}
