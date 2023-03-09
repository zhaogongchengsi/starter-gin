package system

import (
	"gorm.io/gorm"
)

type Languages struct {
	ID        int            `gorm:"primarykey" json:"id"`
	DeletedAt gorm.DeletedAt `json:"-"`
	Name      string         `json:"name" gorm:"index;comment:语言的名字"`
	Value     string         `json:"value" grom:"comment:语言的关键词 根据这个切换语言"`
	Languages []Language     `json:"languages" grom:"many2many:language_Language;comment:关键词"`
}

type Language struct {
	DeletedAt   gorm.DeletedAt
	ID          int         `gorm:"primarykey" json:"id"`
	Key         string      `json:"key" grom:"comment:语言的关键词 根据这个切换语言"`
	Value       string      `json:"value" grom:"comment:词条"`
	LanguagesID int         `json:"languages_id"`
	Languagess  []Languages `gorm:"many2many:language_Language;"`
}

type LanguageKeys struct {
	LanguagesID int `gorm:"primaryKey"`
	LanguageID  int `gorm:"primaryKey"`
	DeletedAt   gorm.DeletedAt
}
