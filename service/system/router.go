package system

import (
	"github.com/zhaogongchengsi/starter-gin/global"
	"github.com/zhaogongchengsi/starter-gin/module"
)

type SysRouter struct {
	Pid       int    `json:"pid"`
	Component string `json:"component" binding:"required"`
	Path      string `json:"path" binding:"required"` // 路由path
	Name      string `json:"name" binding:"required"` // 路由name
	Title     string `json:"title"`
	IsMenu    bool   `json:"isMenu"`
	Icon      string `json:"icon"`
	Auth      bool   `json:"auth"`
	KeepAlive bool   `json:"keepAlive"`
	Hidden    bool   `json:"hidden"` // 是否在列表隐藏
	Sort      int    `json:"sort"`
}

func NewSysRouter() *SysRouter {
	return &SysRouter{}
}

func (sr *SysRouter) GetAllRouters() ([]module.RouterRecord, string, error) {
	rc := new(module.RouterRecord)
	rs, err := rc.GetRouters(global.Db)
	if err != nil {
		return nil, "查询失败", err
	}
	return rs, "查询成功", nil
}
func (sr *SysRouter) CreateRouter() (string, error) {
	r := module.RouterRecord{
		Meta: module.RouterMeTa{
			Title:     sr.Title,
			IsMenu:    sr.IsMenu,
			Icon:      sr.Icon,
			Auth:      sr.Auth,
			KeepAlive: sr.KeepAlive,
		},
		Pid:       sr.Pid,
		Component: sr.Component,
		Path:      sr.Path,
		Name:      sr.Name,
		Hidden:    sr.Hidden,
		Sort:      sr.Sort,
	}

	err := r.CreateRouter(global.Db)

	if err != nil {
		return "创建失败", err
	}

	return "创建成功", err
}

func (sr *SysRouter) DeleteRouter(id int) error {
	r := module.RouterRecord{}
	err := r.DeleteRouterRecord(global.Db, id)
	if err != nil {
		return err
	}
	return nil
}
