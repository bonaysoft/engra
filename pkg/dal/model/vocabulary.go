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
	Phonetic  string `json:"phonetic" csv:"phonetic"`
	Meaning   string `json:"meaning" csv:"meaning"`
	Exchange  string `json:"exchange" csv:"exchange"`
	Tags      string `json:"tags" csv:"tags"`
	Roots     string `json:"roots" csv:"roots"`
	Status    string `json:"status" csv:"status"`
	Parts     string `json:"parts" csv:"-"`
	Intro     string `json:"intro" csv:"-"`
	Etymology string `json:"etymology" csv:"-"`
	NoRoot    bool   `json:"no_root" csv:"-"`
}

func (v *Vocabulary) Row() []string {
	root := "-"
	if len(v.Roots) > 0 {
		roots := strings.Split(v.Roots, ",")
		roots = lo.Map(roots, func(item string, index int) string {
			return fmt.Sprintf("[%s](roots/%s.yml)", item, item)
		})
		root = strings.Join(roots, ",")
	}

	return []string{v.Name, v.Tags, root, v.Status}
}
