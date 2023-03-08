package system

import (
	"mime/multipart"

	"github.com/server-gin/global"
	"github.com/server-gin/modules/system"
	"github.com/server-gin/utils"
)

type UpdateFileInfo struct {
	FileName   string                `json:"fileName"`
	FileHeader *multipart.FileHeader `json:"-"`
}

func (u *UpdateFileInfo) SaveFile() (system.File, error) {

	fn, err := utils.SaveFileHeader(u.FileHeader, global.AppConfig.Server.UploadDir)
	if err != nil {
		return system.File{}, err
	}

	return system.File{FileName: fn}, nil
}
