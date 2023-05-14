package ra

import (
	"testing"

	"github.com/bonaysoft/engra/pkg/dal/model"
	"github.com/stretchr/testify/assert"
)

var wants = []model.Vocabulary{
	// 短单词
	{Name: "as"},
	{Name: "or"},

	// 组合单词
	// {Name: "serviceman", Words: []string{"service", "main"}},

	// 前缀+根
	{Name: "exit", Prefix: "ex", Roots: "it"},
	{Name: "interact", Prefix: "inter", Roots: "act"},
	{Name: "observe", Prefix: "ob", Roots: "serv"},

	// 根+后缀
	{Name: "hydrate", Roots: "hydr", Suffix: "ate"},
	{Name: "service", Roots: "serv", Suffix: "ice"},
	{Name: "stalling", Roots: "stall", Suffix: "ing"},

	// 前缀+根+后缀
	{Name: "detection", Prefix: "de", Roots: "tect", Suffix: "ion"},
	{Name: "protestor", Prefix: "pro", Roots: "test", Suffix: "or"},
	{Name: "protestation", Prefix: "pro", Roots: "test", Suffix: "ation"},
	{Name: "protesting", Prefix: "pro", Roots: "test", Suffix: "ing"},

	// 根+根 TODO
	// {Name: "technocracy", Root: "tect", Suffix: "ion"},
}

func TestExtract(t *testing.T) {
	for _, want := range wants {
		t.Run(want.Name, func(t *testing.T) {
			w := Extract(want.Name)
			assert.Equal(t, w, want)
		})
	}
}
