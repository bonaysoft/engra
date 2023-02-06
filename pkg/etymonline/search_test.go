package etymonline

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

var words = map[string][]string{
	"interact": {"inter-", "act"},
}

func TestSearch(t *testing.T) {
	for word, want := range words {
		t.Run(word, func(t *testing.T) {
			v := Search("interact")
			assert.Equal(t, want, v)
		})
	}
}
