package module

import (
	"gorm.io/gorm"
)

type Languages struct {
	ID        int            `gorm:"primarykey" json:"id"`
	DeletedAt gorm.DeletedAt `json:"-"`
	Name      string         `json:"name" gorm:"index;comment:语言的名字"`
	Value     string         `json:"value" grom:"comment:语言的关键词 前端会根据这个切换语言"`
	Languages []Language     `json:"languages" grom:"comment:关键词"`
}

type Language struct {
	DeletedAt   gorm.DeletedAt
	ID          int    `gorm:"primarykey" json:"id"`
	Key         string `json:"key" grom:"comment:语言的关键词 根据这个切换语言"`
	Value       string `json:"value" grom:"comment:词条"`
	LanguagesID int    `json:"languages_id"`
}

func (la *Language) Transform() map[string]string {
	return map[string]string{
		la.Key: la.Value,
	}
}

func (L *Languages) GetLanguagess(db *gorm.DB) ([]Languages, error) {
	var langs []Languages
	err := db.Model(&Languages{}).Preload("Languages").Find(&langs).Error
	return langs, err
}
