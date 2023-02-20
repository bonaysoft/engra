package model

import (
	"fmt"
	"strings"

	"github.com/samber/lo"
)

// Vocabulary 保存单词的词根词缀
type Vocabulary struct {
	Id        int    `json:"id" csv:"-"`
	Name      string `json:"name" csv:"name"`
	Tag       string `json:"tag" csv:"tag"`
	Root      string `json:"root" csv:"root"`
	Status    string `json:"status" csv:"status"`
	Parts     string `json:"parts" csv:"-"`
	Intro     string `json:"intro" csv:"-"`
	Etymology string `json:"etymology" csv:"-"`
	NoRoot    bool   `json:"no_root" csv:"-"`
}

func (v *Vocabulary) Row() []string {
	root := "-"
	if len(v.Root) > 0 {
		roots := strings.Split(v.Root, ",")
		roots = lo.Map(roots, func(item string, index int) string {
			return fmt.Sprintf("[%s](roots/%s.yml)", item, item)
		})
		root = strings.Join(roots, ",")
	}

	return []string{v.Name, v.Tag, root, v.Status}
}
