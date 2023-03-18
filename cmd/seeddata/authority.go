package seeddata

import (
	"github.com/server-gin/modules/system"
	"gorm.io/gorm"
)

func CreateAuthorityTable(db *gorm.DB) error {
	return db.AutoMigrate(system.Authority{})
}
func CrateAuthouitySeedData(db *gorm.DB) error {
	return nil
}
