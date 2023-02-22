package client

import (
	"context"
	"net/http"

	"github.com/Khan/genqlient/graphql"
)

type Client struct {
	graphql.Client
}

func (c *Client) WithContext(ctx context.Context) context.Context {
	return context.WithValue(ctx, "client", c)
}

type Context struct {
	Client
}

func NewClient(endpoint string) *Client {
	return &Client{
		Client: graphql.NewClient(endpoint, http.DefaultClient),
	}
}
