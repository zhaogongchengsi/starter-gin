package module

type File struct {
	BaseMode
	FileName string `json:"fileName"`
}

func (File) TableName() string {
	return "file"
}
