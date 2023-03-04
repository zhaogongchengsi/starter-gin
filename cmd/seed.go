package cmd

import (
	"fmt"

	"github.com/server-gin/modules/system"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func seed(dsn string) error {

	db, err := gorm.Open(mysql.New(mysql.Config{
		DSN:                       dsn,
		DefaultStringSize:         256,   // string 类型字段的默认长度
		DisableDatetimePrecision:  true,  // 禁用 datetime 精度，MySQL 5.6 之前的数据库不支持
		DontSupportRenameIndex:    true,  // 重命名索引时采用删除并新建的方式，MySQL 5.7 之前的数据库和 MariaDB 不支持重命名索引
		DontSupportRenameColumn:   true,  // 用 `change` 重命名列，MySQL 8 之前的数据库和 MariaDB 不支持重命名列
		SkipInitializeWithVersion: false, // 根据当前 MySQL 版本自动配置
	}), &gorm.Config{})

	if err != nil {
		return fmt.Errorf("seed Error: Database connection failed, please check %s and try again", dsn)
	}

	var user system.User

	err = db.AutoMigrate(user)

	if err != nil {
		return fmt.Errorf("seed Error: database initialization failed : %v", err)
	}

	newUser := system.NewUser("admin", "12345", "18312391231", "管理员", "zzh123123123@qq.com")

	db.Create(newUser)

	return nil
}
