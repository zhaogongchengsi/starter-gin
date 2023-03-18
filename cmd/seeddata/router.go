package seeddata

import (
	"github.com/server-gin/modules/system"
	"gorm.io/gorm"
)

func CreateRouter(db *gorm.DB) error {
	return db.AutoMigrate(system.RouterRecord{})
}

var Routers = []system.RouterRecord{
	{
		Pid:       0,
		Path:      "dashboard",
		Component: "/views/dashboard",
		Name:      "dashboard",
		Meta: system.RouterMeTa{
			Title:  "router.title.dashboard",
			Auth:   false,
			IsMenu: true,
			Icon:   "icon-dashboard",
		},
	},
	{
		Pid:       1,
		Path:      "workplace",
		Component: "/views/dashboard/workplace.vue",
		Name:      "workplace",
		Meta: system.RouterMeTa{
			Title:  "router.title.workbench",
			Auth:   false,
			IsMenu: true,
			Icon:   "icon-common",
		},
	},
	{
		Pid:       0,
		Path:      "notComponent",
		Component: "/views/notComponent",
		Name:      "notExist",
		Meta: system.RouterMeTa{
			Title:  "router.title.abnormal",
			Auth:   false,
			IsMenu: true,
			Icon:   "icon-exclamation-polygon-fill",
		},
	},
	{
		Pid:       0,
		Path:      "utils",
		Component: "/views/utils",
		Name:      "utils",
		Meta: system.RouterMeTa{
			Title:  "router.title.toolLibrary",
			Auth:   false,
			IsMenu: true,
			Icon:   "icon-calendar",
		},
	},
	{
		Pid:       4,
		Path:      "fileSplit",
		Component: "/views/utils/FileSplit",
		Name:      "fileSplit",
		Meta: system.RouterMeTa{
			Title:  "router.title.fileSplitting",
			Auth:   false,
			IsMenu: true,
			Icon:   "icon-file",
		},
	},
}

func CrateRouterSeedData(db *gorm.DB) error {
	return db.Model(system.RouterRecord{}).Create(Routers).Error
}
