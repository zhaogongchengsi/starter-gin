package seeddata

import (
	"github.com/server-gin/module"
	"gorm.io/gorm"
)

func CreateLanguages(db *gorm.DB) error {
	return db.AutoMigrate(&module.Languages{}, &module.Language{})
}
func CreateLanguage(db *gorm.DB) error {
	return db.AutoMigrate(&module.Language{})
}

var Languages = []module.Languages{
	{
		Name:  "英文",
		Value: "en",
		Languages: []module.Language{
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
		Languages: []module.Language{
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
}

func CrateLanguagesSeedData(db *gorm.DB) error {
	return db.Model(module.Languages{}).Create(Languages).Error
}
