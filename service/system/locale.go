package system

import (
	"github.com/server-gin/global"
	"github.com/server-gin/modules/system"
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
	lang := system.Languages{}
	langs, err := lang.GetLanguagess(global.Db)

	if err != nil {
		return messages, err
	}

	for _, l := range langs {
		lang := NewLanguages(l.Value, l.Name)
		for _, l2 := range l.Languages {
			lang.Languages[l2.Key] = l2.Value
		}
		messages = append(messages, *lang)
	}

	return messages, err
}
