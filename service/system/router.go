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

func (sr *SysRouter) GetUserRouters(phone, password string) ([]module.RouterRecord, string, error) {
	user := module.NewFindUser(phone, password)
	list, err := user.GetAuthoritysByPhone(global.Db)
	if err != nil {
		return []module.RouterRecord{}, "获取路由失败", err
	}
	routers := []module.RouterRecord{}

	// 权限内所有的路由合并起来
	for _, authority := range list {
		routers = append(routers, authority.RouterRecords...)
	}

	return routers, "获取成功", nil
}
