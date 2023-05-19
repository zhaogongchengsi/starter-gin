package seeddata

import (
	"github.com/zhaogongchengsi/starter-gin/module"
	"gorm.io/gorm"
)

func CreateRouter(db *gorm.DB) error {
	return db.AutoMigrate(module.RouterRecord{})
}

func CrateRouterSeedData(db *gorm.DB) error {
	return db.Model(module.RouterRecord{}).Create(Routers).Error
}

var Routers = []module.RouterRecord{
	// 工作台路由
	{
		Pid:       0,
		Path:      "dashboard",
		Component: "/views/dashboard",
		Name:      "dashboard",
		Meta: module.RouterMeTa{
			Title:  "router.title.dashboard",
			Auth:   false,
			IsMenu: true,
			Icon:   "icon-dashboard",
		},
		Sort: 0,
	},
	{
		Pid:       1,
		Path:      "workplace",
		Component: "/views/dashboard/workplace.vue",
		Name:      "workplace",
		Meta: module.RouterMeTa{
			Title:  "router.title.workbench",
			Auth:   false,
			IsMenu: true,
			Icon:   "icon-common",
		},
		Sort: 1,
	},
	// 子仓库
	{
		Pid:       0,
		Path:      "rules",
		Component: "/views/rules/index.vue",
		Name:      "roleAdmin",
		Meta: module.RouterMeTa{
			Title:  "router.title.roles",
			Auth:   true,
			IsMenu: true,
			Icon:   "icon-common",
		},
		Sort: 2,
	},
}
