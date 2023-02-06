package etymonline

import (
	"fmt"
	"strings"

	"github.com/antchfx/htmlquery"
	"github.com/samber/lo"
	"golang.org/x/net/html"
)

func Search(word string) []string {
	doc, err := htmlquery.LoadURL("https://www.etymonline.com/search?q=" + word)
	if err != nil {
		panic(err)
	}
	// Find all news item.
	list, err := htmlquery.QueryAll(doc, "//div/div/a")
	if err != nil {
		panic(err)
	}

	for _, n := range list {
		a := htmlquery.FindOne(n, "//a[@title]")
		title := htmlquery.SelectAttr(a, "title")
		if strings.TrimPrefix(title, "Origin and meaning of ") != word {
			continue
		}

		fmt.Println(title)
		b2, err := htmlquery.QueryAll(n, "//a[@title]/../div/section/p/a")
		if err != nil {
			return nil
		}

		return lo.Map(b2, func(item *html.Node, index int) string { return htmlquery.InnerText(item) })
	}

	return nil
}
