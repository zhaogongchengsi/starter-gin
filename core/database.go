package core

import (
	"fmt"
	"time"

	"github.com/zhaogongchengsi/starter-gin/config"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/schema"
)

func ConnectDataBaseServer(cfg *config.Config) (*gorm.DB, error) {

	config := cfg.DataBase

	dns := fmt.Sprintf(`%s:%s@tcp(%s:%v)/%s?charset=%s&parseTime=True&loc=Local`,
		config.Username,
		config.Password,
		config.Url,
		config.Port,
		config.DbName,
		config.Charset,
	)

	DB, err := gorm.Open(mysql.New(mysql.Config{
		DSN:                       dns,
		DefaultStringSize:         256,   // string 类型字段的默认长度
		DisableDatetimePrecision:  true,  // 禁用 datetime 精度，MySQL 5.6 之前的数据库不支持
		DontSupportRenameIndex:    true,  // 重命名索引时采用删除并新建的方式，MySQL 5.7 之前的数据库和 MariaDB 不支持重命名索引
		DontSupportRenameColumn:   true,  // 用 `change` 重命名列，MySQL 8 之前的数据库和 MariaDB 不支持重命名列
		SkipInitializeWithVersion: false, // 根据当前 MySQL 版本自动配置
	}), &gorm.Config{
		NamingStrategy: schema.NamingStrategy{
			TablePrefix:   config.TablePrefix,   // 表名前缀，
			SingularTable: config.SingularTable, // 使用单数表名，启用该选项后，`User` 表将是`user
		},
		DisableForeignKeyConstraintWhenMigrating: true, // 不使用物理外键
	})

	if err != nil {
		return nil, err
	}

	sqlDB, err := DB.DB()

	if err != nil {
		return nil, err
	}

	// SetMaxIdleConns 用于设置连接池中空闲连接的最大数量。
	sqlDB.SetMaxIdleConns(config.MaxIdleConns)
	//// SetMaxOpenConns 设置打开数据库连接的最大数量。
	sqlDB.SetMaxOpenConns(config.MaxOpenConns)
	// SetConnMaxLifetime 设置了连接可复用的最大时间。
	sqlDB.SetConnMaxLifetime(time.Hour)

	return DB, nil

}
