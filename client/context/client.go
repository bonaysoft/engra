package context

import (
	"context"
	"fmt"

	"github.com/Khan/genqlient/graphql"
)

func GetClient(ctx context.Context) (graphql.Client, error) {
	v, ok := ctx.Value("client").(graphql.Client)
	if !ok {
		return nil, fmt.Errorf("not found client")
	}

	return v, nil
}
