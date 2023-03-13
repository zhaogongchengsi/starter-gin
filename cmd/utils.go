package cmd

import (
	"fmt"

	uuid "github.com/satori/go.uuid"
	"github.com/server-gin/core"
	"github.com/server-gin/global"
	"github.com/server-gin/modules/system"
	"gorm.io/gorm"
)

var moduleMap = map[string]any{
	"user":      system.User{},
	"file":      system.File{},
	"router":    system.RouterRecord{},
	"languages": system.Languages{},
	"language":  system.Language{},
}

var moduleSeedMap = map[string]any{
	"user": []system.User{
		{
			Phone:    "12312312312",
			Password: system.CreatePassworld("123456"),
			UUID:     uuid.NewV4(),
			UserName: "admin",
			NickName: "超级管理员",
		},
	},
	"router": []system.RouterRecord{
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
	},
	"languages": []system.Languages{
		{
			Name:  "英文",
			Value: "en",
			Languages: []system.Language{
				{Key: "hello", Value: "hello"},
				{Key: "loginpage.from.place.account", Value: "Please enter your phone number"},
				{Key: "loginpage.from.place.password", Value: "Please enter your password"},
				{Key: "loginpage.from.place.verifi", Value: "Please enter your verification code"},
				{Key: "loginpage.from.account", Value: "Phone"},
				{Key: "loginpage.from.password", Value: "Password"},
				{Key: "loginpage.from.verifi", Value: "Verification code"},
				{Key: "loginpage.from.login", Value: "login"},
				{Key: "header.right.logout", Value: "logout"},
				{Key: "header.right.setting", Value: "setting"},
				{Key: "outer.title.home", Value: "Home"},
				{Key: "router.title.dashboard", Value: "Dashboard"},
				{Key: "router.title.workbench", Value: "Workbench"},
				{Key: "router.title.abnormal", Value: "Abnormal"},
				{Key: "router.title.toolLibrary", Value: "Tool Library"},
				{Key: "router.title.fileSplitting", Value: "file splitting"},
			},
		},
		{
			Name:  "中文",
			Value: "cn",
			Languages: []system.Language{
				{Key: "hello", Value: "你好"},
				{Key: "loginpage.from.place.account", Value: "请输入手机号"},
				{Key: "loginpage.from.place.password", Value: "请输入密码"},
				{Key: "loginpage.from.place.verifi", Value: "请输入验证码"},
				{Key: "loginpage.from.account", Value: "手机号"},
				{Key: "loginpage.from.password", Value: "密码"},
				{Key: "loginpage.from.verifi", Value: "验证码"},
				{Key: "loginpage.from.login", Value: "登陆"},
				{Key: "header.right.logout", Value: "注销登陆"},
				{Key: "header.right.setting", Value: "系统设置"},
				{Key: "outer.title.home", Value: "首页"},
				{Key: "router.title.dashboard", Value: "仪表盘"},
				{Key: "router.title.workbench", Value: "工作台"},
				{Key: "router.title.abnormal", Value: "异常组件"},
				{Key: "router.title.toolLibrary", Value: "工具库"},
				{Key: "router.title.fileSplitting", Value: "文件切分"},
			},
		},
	},
}

func ConnDb(file, typ, name string) (*gorm.DB, error) {

	conf, err := global.ReadConfig(file, typ, name)

	if err != nil {
		return &gorm.DB{}, fmt.Errorf("seed Error: The specified parameters are wrong, and the database configuration cannot be obtained. %s %s %v", file, typ, err)
	}

	db, err := core.ConnectDataBaseServer(conf)

	if err != nil {
		return &gorm.DB{}, fmt.Errorf("seed Error: Database connection failed, please check %s and try again", err)
	}

	return db, nil
}
