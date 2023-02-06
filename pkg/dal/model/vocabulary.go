package model

// Vocabulary 保存单词的词根词缀
type Vocabulary struct {
	Id        int    `json:"id"`
	Name      string `json:"name"`
	Tag       string `json:"tag"`
	Parts     string `json:"parts"`
	Intro     string `json:"intro"`
	Etymology string `json:"etymology"`
}
