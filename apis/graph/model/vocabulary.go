package model

import (
	"strconv"
)

func NewVocabulary() *Vocabulary {
	return &Vocabulary{}
}

func (n *Vocabulary) Status() string {
	var score int
	if n.Phonetic != "" {
		score++
	}
	if n.Meaning != "" {
		score++
	}
	if n.Mnemonic != "" {
		score++
	}
	if len(n.Constitute) > 0 {
		score++
	}
	if len(n.Children) > 0 {
		score++
	}

	if score > 0 {
		return strconv.Itoa(score)
	}

	return "-"
}

func (n *Vocabulary) Find(word string) (*Vocabulary, bool) {
	if n.Name == word {
		return n, true
	}

	for _, child := range n.Children {
		v, ok := child.Find(word)
		if ok {
			return v, true
		}
	}

	return nil, false
}
