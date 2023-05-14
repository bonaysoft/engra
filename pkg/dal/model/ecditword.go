package model

type EcDictWord struct {
	Word        string
	Phonetic    string
	Definition  string
	Translation string
	Tag         string
	Exchange    string
}

func (m *EcDictWord) TableName() string {
	return "stardict"
}
