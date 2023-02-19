package model

import "github.com/samber/lo"

func NewVocabulary() *Vocabulary {
	return &Vocabulary{}
}

func (n *Vocabulary) Exist(word string) bool {
	if n.Name == word {
		return true
	}

	if len(n.Children) == 0 {
		return false
	}

	_, ok := lo.Find(n.Children, func(item *Vocabulary) bool { return item.Exist(word) })
	return ok
}