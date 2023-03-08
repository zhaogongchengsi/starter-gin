package system

import common "github.com/server-gin/modules"

type File struct {
	common.BaseMode
	FileName string `json:"fileName"`
}
