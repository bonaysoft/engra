package dict

import (
	"embed"
	"fmt"
	"io/fs"

	"github.com/bonaysoft/engra/apis/graph/model"
)

//go:embed roots
var rootsFs embed.FS

type Roots struct {
	trees []*WordRoot
}

func NewRoots() (*Roots, error) {
	var trees = make([]*WordRoot, 0)
	fn := func(path string, d fs.DirEntry, err error) error {
		if d.IsDir() {
			return nil
		}

		tree, err2 := NewWordRoot(path)
		if err2 != nil {
			return err2
		}

		trees = append(trees, tree)
		return nil
	}

	if err := fs.WalkDir(rootsFs, ".", fn); err != nil {
		return nil, err
	}
	return &Roots{trees: trees}, nil
}

func (d *Roots) Find(wordStr string) (*model.Vocabulary, *WordRoot, error) {
	for _, tree := range d.trees {
		v, ok := tree.Find(wordStr)
		if ok {
			return v, tree, nil
		}
	}

	return nil, nil, fmt.Errorf("not found")
}
