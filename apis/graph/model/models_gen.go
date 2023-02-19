// Code generated by github.com/99designs/gqlgen, DO NOT EDIT.

package model

type Vocabulary struct {
	Name       string        `json:"name"`
	Phonetic   string        `json:"phonetic" yaml:"phonetic,omitempty"`
	Mnemonic   string        `json:"mnemonic" yaml:"mnemonic,omitempty"`
	Constitute []string      `json:"constitute" yaml:"constitute,omitempty"`
	Meaning    string        `json:"meaning" yaml:"meaning,omitempty"`
	Tags       []string      `json:"tags" yaml:"tags,omitempty"`
	Children   []*Vocabulary `json:"children" yaml:"children,omitempty"`
}
