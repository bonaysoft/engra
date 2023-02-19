package dict

import (
	"fmt"
	"io/fs"
	"path/filepath"
	"strings"

	"github.com/samber/lo"
)

type Dict struct {
	trees []*WordRoot
}

func NewDict() (*Dict, error) {
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

	if err := filepath.WalkDir("dicts", fn); err != nil {
		return nil, err
	}
	return &Dict{trees: trees}, nil
}

func (d *Dict) Find(word string) (*WordRoot, error) {
	v, ok := lo.Find(d.trees, func(n *WordRoot) bool { return n.Exist(word) })
	if !ok {
		return nil, fmt.Errorf("[%s] not found at tree", word)
	}

	return v, nil
}
