package model

// Vocabulary 保存单词的词根词缀
type Vocabulary struct {
	Id     int      `json:"id"`
	Name   string   `json:"name"`
	Tag    string   `json:"tag"`
	Words  []string `json:"words" gorm:"-"`
	Root   string   `json:"root"`
	Prefix string   `json:"prefix"`
	Suffix string   `json:"suffix"`
}
