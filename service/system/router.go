package system

import (
	"github.com/zhaogongchengsi/starter-gin/global"
	"github.com/zhaogongchengsi/starter-gin/module"
)

type SysRouter struct {
}

func NewSysRouter() *SysRouter {
	return &SysRouter{}
}

func (sr *SysRouter) GetUserRouters() ([]module.RouterRecord, string, error) {
	rc := new(module.RouterRecord)

	rs, err := rc.GetRouters(global.Db)

	if err != nil {
		return nil, "查询失败", err
	}
	return rs, "查询成功", nil
}
