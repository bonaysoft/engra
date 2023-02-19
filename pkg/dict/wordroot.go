package dict

import (
	"os"

	"github.com/bitfield/script"
	"github.com/bonaysoft/engra/apis/graph/model"
	"gopkg.in/yaml.v3"
)

type WordRoot struct {
	*model.Vocabulary

	path string
}

func NewWordRoot(name string) (*WordRoot, error) {
	path := "./dicts/" + name + ".yml"
	content, err := script.File(path).Bytes()
	if err != nil {
		return nil, err
	}

	tree := model.NewVocabulary()
	if err := yaml.Unmarshal(content, tree); err != nil {
		return nil, err
	}

	return &WordRoot{
		Vocabulary: tree,
		path:       path,
	}, nil
}

func (wr *WordRoot) Save() error {
	f, err := os.OpenFile(wr.path, os.O_WRONLY|os.O_TRUNC, 0666)
	if err != nil {
		return err
	}

	ye := yaml.NewEncoder(f)
	ye.SetIndent(2)
	return ye.Encode(wr.Vocabulary)
}
