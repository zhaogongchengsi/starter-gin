package cmd

import (
	"fmt"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Phone    string
	UserName string
	NickName string
	Password string
}

func newUser(ph, un, nn, pw string) *User {
	return &User{
		Phone:    ph,
		Password: pw,
		UserName: un,
		NickName: nn,
	}
}

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

	err = db.AutoMigrate(&User{})

	if err != nil {
		return fmt.Errorf("seed Error: database initialization failed")
	}

	db.Create(newUser("12312312312", "user", "admin", "123abc"))

	return nil
}
