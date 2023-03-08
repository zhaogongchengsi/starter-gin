package system

type Languages struct {
	ID        uint       `gorm:"primarykey" json:"id"`
	Name      string     `json:"name" gorm:"index;comment:语言的名字"`
	Value     string     `json:"value" grom:"comment:语言的关键词 根据这个切换语言"`
	Languages []Language `json:"Languages" grom:"foreignKey:Key;comment:关键词"`
}

type Language struct {
	ID          uint   `gorm:"primarykey" json:"id"`
	Key         string `json:"key" grom:"comment:语言的关键词 根据这个切换语言"`
	Value       string `json:"value" grom:"comment:词条"`
	LanguagesID uint
}
