package ra

import (
	"testing"

	"github.com/bonaysoft/engra/pkg/dal/model"
	"github.com/go-playground/assert/v2"
)

var wants = []model.Vocabulary{
	// 短单词
	{Name: "as"},
	{Name: "or"},

	// 组合单词
	// {Name: "serviceman", Words: []string{"service", "main"}},

	// 前缀+根
	{Name: "exit", Prefix: "ex", Root: "it"},
	{Name: "interact", Prefix: "inter", Root: "act"},
	{Name: "observe", Prefix: "ob", Root: "serv"},

	// 根+后缀
	{Name: "hydrate", Root: "hydr", Suffix: "ate"},
	{Name: "service", Root: "serv", Suffix: "ice"},
	{Name: "stalling", Root: "stall", Suffix: "ing"},

	// 前缀+根+后缀
	{Name: "detection", Prefix: "de", Root: "tect", Suffix: "ion"},
	{Name: "protestor", Prefix: "pro", Root: "test", Suffix: "or"},
	{Name: "protestation", Prefix: "pro", Root: "test", Suffix: "ation"},
	{Name: "protesting", Prefix: "pro", Root: "test", Suffix: "ing"},

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
