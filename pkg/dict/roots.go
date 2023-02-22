package dict

import (
	"fmt"
	"io/fs"
	"path/filepath"
	"strings"

	"github.com/bonaysoft/engra/apis/graph/model"
)

type Roots struct {
	trees []*WordRoot
}

func NewRoots() (*Roots, error) {
	var trees = make([]*WordRoot, 0)
	fn := func(path string, d fs.DirEntry, err error) error {
		if d.IsDir() {
			return nil
		}

		tree, err2 := NewWordRoot(strings.TrimSuffix(d.Name(), ".yml"))
		if err2 != nil {
			return err2
		}

		trees = append(trees, tree)
		return nil
	}

	if err := filepath.WalkDir("dict/roots", fn); err != nil {
		return nil, err
	}
	return &Roots{trees: trees}, nil
}

func (d *Roots) Find(wordStr string) (*model.Vocabulary, *model.Vocabulary, error) {
	var word, root *model.Vocabulary
	for _, tree := range d.trees {
		v, ok := tree.Find(wordStr)
		if ok {
			word = v
			root = tree.Vocabulary
			break
		}
	}

	if word == nil {
		return nil, nil, fmt.Errorf("not found")
	}

	return word, root, nil
}
