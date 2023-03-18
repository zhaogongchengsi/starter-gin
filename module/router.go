package module

type RouterMeTa struct {
	Title     string `json:"title" gorm:"comment:路由标题"`
	IsMenu    bool   `json:"isMenu" gorm:"comment:是否需要在菜单栏显示"`
	Icon      string `json:"icon" gorm:"comment:路由的icon图标"`
	Auth      bool   `json:"auth" gorm:"comment:是否需要鉴权"`
	KeepAlive bool   `json:"keepAlive" gorm:"comment:是否缓存"`
}

type RouterRecord struct {
	BaseMode
	Meta       RouterMeTa  `json:"meta" gorm:"embedded;comment:路由元信息"`
	Pid        int         `json:"pid" gorm:"index;comment:根据此pid判断父级路由是谁"`
	Component  string      `json:"component" gorm:"comment:路由组件"`
	Path       string      `json:"path" gorm:"comment:路由path"`    // 路由path
	Name       string      `json:"name" gorm:"comment:路由name"`    // 路由name
	Hidden     bool        `json:"hidden" gorm:"comment:是否在列表隐藏"` // 是否在列表隐藏
	Sort       int         `json:"sort" gorm:"comment:排序标记"`
	Authoritys []Authority `json:"authoritys" gorm:"many2many:authority_routers;"`
}
