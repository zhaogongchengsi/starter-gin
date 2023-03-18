package system

import (
	"github.com/server-gin/module"
	"mime/multipart"

	"github.com/server-gin/global"
	"github.com/server-gin/utils"
)

type UpdateFileInfo struct {
	FileName   string                `json:"fileName"`
	FileHeader *multipart.FileHeader `json:"-"`
}

func (u *UpdateFileInfo) SaveFile() (module.File, error) {

	fn, err := utils.SaveFileHeader(u.FileHeader, global.AppConfig.Server.UploadDir)
	if err != nil {
		return module.File{}, err
	}

	return module.File{FileName: fn}, nil
}
