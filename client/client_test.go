package client

import (
	"context"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestName(t *testing.T) {
	ctx := context.Background()
	c := NewClient("http://localhost:8081/query")
	ctx = c.WithContext(ctx)
	resp, err := Find(ctx, "animal")
	assert.NoError(t, err)
	assert.Equal(t, "animal", resp.GetVocabulary().Self.Name)
}
