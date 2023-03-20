package system

import (
	"github.com/zhaogongchengsi/starter-gin/global"
	"github.com/zhaogongchengsi/starter-gin/module"
)

type Messages struct{}

type Languages struct {
	Name      string            `json:"name"`
	Label     string            `json:"label"`
	Languages map[string]string `json:"languages"`
}

func NewLanguages(name, label string) *Languages {
	return &Languages{
		Name:      name,
		Label:     label,
		Languages: make(map[string]string),
	}
}

func (L *Messages) GetMessages() ([]Languages, error) {

	var messages []Languages
	lang := module.Languages{}
	langs, err := lang.GetLanguagess(global.Db)

	if err != nil {
		return messages, err
	}

	messages = make([]Languages, len(langs))
	for i, l := range langs {
		lang := NewLanguages(l.Value, l.Name)
		for _, l2 := range l.Languages {
			lang.Languages[l2.Key] = l2.Value
		}
		messages[i] = *lang
		// messages = append(messages, *lang)
	}

	return messages, err
}
