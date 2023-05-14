package dict

import (
	"fmt"
	"testing"

	"github.com/bonaysoft/engra/pkg/dal/model"
	"github.com/samber/lo"
	"github.com/stretchr/testify/assert"
)

func TestGlossaries2Words(t *testing.T) {
	assert.NoError(t, Glossaries2Words())
}

func TestGetWords(t *testing.T) {
	v, ok := lo.Find(GetWords(), func(item model.Vocabulary) bool {
		return item.Name == "eighteen"
	})
	assert.True(t, ok)
	assert.NotEmpty(t, v.Phonetic)
	assert.NotEmpty(t, v.Meaning)
	assert.NotEmpty(t, v.Tags)
	fmt.Println(v)
}
