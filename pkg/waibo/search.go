package waibo

import (
	"fmt"
	"strings"

	"github.com/antchfx/htmlquery"
)

func Search2(word string) ([]string, string) {
	doc, err := htmlquery.LoadURL("https://www.waibo.wang/r/" + word)
	if err != nil {
		panic(err)
	}
	// Find all news item.

	structTip := htmlquery.FindOne(doc, "//p[@class='struct-tip']")
	if structTip == nil {
		return nil, ""
	}

	var story string
	storyTip := htmlquery.FindOne(doc, "//div[@class='story-tip']")
	if storyTip != nil {
		story = htmlquery.OutputHTML(storyTip, false)
	}

	v := htmlquery.InnerText(structTip)
	s1, s2 := v, ""
	spIdx := strings.IndexRune(v, 8594)
	if spIdx > -1 {
		s1 = v[:spIdx]
		s2 = v[spIdx:]
	}

	return extract(s1), fmt.Sprintf("%s ----> %s", s2, story)
}

func Search(word string) []string {
	return nil
}

func extract(text string) []string {
	var result []string
	var tmp string
	for _, v := range text {
		if (v >= 97 && v <= 122) || v == 45 {
			tmp += string(v)
		}

		if len(tmp) > 1 && (v == 45 || v > 122) {
			result = append(result, tmp)
			tmp = ""
		}
	}
	return result
}
