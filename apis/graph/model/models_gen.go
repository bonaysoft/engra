// Code generated by github.com/99designs/gqlgen, DO NOT EDIT.

package model

type Vocabulary struct {
	Name     string        `json:"name"`
	Children []*Vocabulary `json:"children"`
}
