package system

import (
	"github.com/zhaogongchengsi/starter-gin/module"
	"mime/multipart"

	"github.com/zhaogongchengsi/starter-gin/global"
	"github.com/zhaogongchengsi/starter-gin/utils"
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
