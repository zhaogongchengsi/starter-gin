package core

import (
	"errors"
	"fmt"
	"time"

	"github.com/zhaogongchengsi/starter-gin/config"

	"gorm.io/driver/mysql"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/schema"
)

var (
	ErrDbTypeExit = errors.New("err: database type does not exist")
)

func ConnectDataBaseServer(cfg *config.Config) (*gorm.DB, error) {

	dbConfig := cfg.DataBase

	connect, err := GetDBConnect(&dbConfig)

	if err != nil {
		return nil, err
	}

	DB, err := gorm.Open(connect, &gorm.Config{
		NamingStrategy: schema.NamingStrategy{
			TablePrefix:   dbConfig.TablePrefix,   // 表名前缀，
			SingularTable: dbConfig.SingularTable, // 使用单数表名，启用该选项后，`User` 表将是`user
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
	sqlDB.SetMaxIdleConns(dbConfig.MaxIdleConns)
	//// SetMaxOpenConns 设置打开数据库连接的最大数量。
	sqlDB.SetMaxOpenConns(dbConfig.MaxOpenConns)
	// SetConnMaxLifetime 设置了连接可复用的最大时间。
	sqlDB.SetConnMaxLifetime(time.Hour)

	return DB, nil

}

func GetDBConnect(config *config.DataBase) (gorm.Dialector, error) {
	switch config.DbType {
	case "mysql":
		return CreateMysqlConnect(config), nil
	case "postgres":
		return CreatePostgresConnect(config), nil
		// 后续可以扩展其他类型
	default:
		return nil, ErrDbTypeExit
	}
}

func CreateMysqlConnect(config *config.DataBase) gorm.Dialector {
	dns := fmt.Sprintf(`%s:%s@tcp(%s:%v)/%s?charset=%s&parseTime=True&loc=Local`,
		config.Username,
		config.Password,
		config.Url,
		config.Port,
		config.DbName,
		config.Charset,
	)
	return mysql.New(mysql.Config{
		DSN:                       dns,
		DefaultStringSize:         256,   // string 类型字段的默认长度
		DisableDatetimePrecision:  true,  // 禁用 datetime 精度，MySQL 5.6 之前的数据库不支持
		DontSupportRenameIndex:    true,  // 重命名索引时采用删除并新建的方式，MySQL 5.7 之前的数据库和 MariaDB 不支持重命名索引
		DontSupportRenameColumn:   true,  // 用 `change` 重命名列，MySQL 8 之前的数据库和 MariaDB 不支持重命名列
		SkipInitializeWithVersion: false, // 根据当前 MySQL 版本自动配置
	})
}

func CreatePostgresConnect(config *config.DataBase) gorm.Dialector {
	dns := fmt.Sprintf("postgres://%s:%s@%s:%v/%s",
		config.Username,
		config.Password,
		config.Url,
		config.Port,
		config.DbName,
	)

	return postgres.New(postgres.Config{
		DSN:                  dns,
		PreferSimpleProtocol: true, // disables implicit prepared statement usage
	})
}
