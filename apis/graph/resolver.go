package graph

import (
	"context"

	"github.com/bonaysoft/engra/apis/graph/model"
)

// This file will not be regenerated automatically.
//
// It serves as dependency injection for your app, add any dependencies you require here.

type Dict interface {
	Lookup(ctx context.Context, name string) (*model.Vocabulary, error)
	LookupWithRoot(ctx context.Context, name string) (*model.Vocabulary, error)
}

type Resolver struct {
	Dict
}
