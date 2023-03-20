package seeddata

import (
	"github.com/zhaogongchengsi/starter-gin/module"
	"gorm.io/gorm"
)

func CrateFileTable(db *gorm.DB) error {
	return db.AutoMigrate(module.File{})
}

func CrateFileSeedData(db *gorm.DB) error {
	return nil
}
