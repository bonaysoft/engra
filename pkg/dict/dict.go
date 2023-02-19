package dict

import (
	"fmt"
	"io/fs"
	"path/filepath"

	"github.com/bonaysoft/engra/apis/graph/model"
	"github.com/samber/lo"
)

type Dict struct {
	trees []*model.Vocabulary
}

func NewDict() (*Dict, error) {
	var trees = make([]*model.Vocabulary, 0)
	fn := func(path string, d fs.DirEntry, err error) error {
		if d.IsDir() {
			return nil
		}

		tree, err2 := NewWordRoot(path)
		if err2 != nil {
			return err2
		}

		trees = append(trees, tree.Vocabulary)
		return nil
	}

	if err := filepath.WalkDir("dicts", fn); err != nil {
		return nil, err
	}
	return &Dict{trees: trees}, nil
}

func (d *Dict) Find(word string) (*model.Vocabulary, error) {
	v, ok := lo.Find(d.trees, func(n *model.Vocabulary) bool { return n.Exist(word) })
	if !ok {
		return nil, fmt.Errorf("[%s] not found at tree", word)
	}

	return v, nil
}
