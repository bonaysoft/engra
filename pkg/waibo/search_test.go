package waibo

import (
	"fmt"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

// 	{Name: "exit", Prefix: "ex", Root: "it"},
//	{Name: "interact", Prefix: "inter", Root: "act"},
//	{Name: "observe", Prefix: "ob", Root: "serv"},
//
//	// 根+后缀
//	{Name: "hydrate", Root: "hydr", Suffix: "ate"},
//	{Name: "service", Root: "serv", Suffix: "ice"},
//	{Name: "stalling", Root: "stall", Suffix: "ing"},
//
//	// 前缀+根+后缀
//	{Name: "detection", Prefix: "de", Root: "tect", Suffix: "ion"},
//	{Name: "protestor", Prefix: "pro", Root: "test", Suffix: "or"},
//	{Name: "protestation", Prefix: "pro", Root: "test", Suffix: "ation"},
//	{Name: "protesting", Prefix: "pro", Root: "test", Suffix: "ing"},

var words = map[string][]string{
	"exit":     {"ex-", "-it-"},
	"interact": {"inter-", "act"},
	"observe":  {"ob-", "-serv-", "-e"},
	"hydrate":  {"-hydr-", "-ate"},
	"detect":   {"de-", "-tect-"},
	"protest":  {"pro-", "-test-"},
}

func TestSearch(t *testing.T) {
	for word, want := range words {
		t.Run(word, func(t *testing.T) {
			v := Search(word)
			assert.Equal(t, want, v)
		})
	}
}

func TestName(t *testing.T) {
	for _, i := range "→" {
		fmt.Println(i)
	}

	fmt.Println(strings.IndexRune("123→", 8594))
}

func TestFetchTree(t *testing.T) {
	n, err := FetchTree("him")
	fmt.Println(n, err)
}
