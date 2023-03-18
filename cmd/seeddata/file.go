package seeddata

import (
	"github.com/server-gin/modules/system"
	"gorm.io/gorm"
)

func CrateFileTable(db *gorm.DB) error {
	return db.AutoMigrate(system.File{})
}

func CrateFileSeedData(db *gorm.DB) error {
	return nil
}
