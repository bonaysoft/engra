package dict

import (
	_ "embed"

	"github.com/bonaysoft/engra/pkg/dal/model"
	"github.com/gocarina/gocsv"
)

//go:embed words.csv
var wordCsv []byte

func GetWords() []model.Vocabulary {
	words := make([]model.Vocabulary, 0)
	gocsv.UnmarshalBytes(wordCsv, &words)
	return words
}
