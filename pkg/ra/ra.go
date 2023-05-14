package ra

import (
	_ "embed"
	"sort"
	"strings"

	"github.com/bonaysoft/engra/pkg/dal/model"
	"github.com/samber/lo"
)

//go:embed ra.txt
var raBytes []byte

var gRootAffixes []string

func init() {
	gRootAffixes = strings.Split(string(raBytes), "\n")
}

func SetupRootAffixes(rootAffixes []string) {
	gRootAffixes = rootAffixes
}

func Extract(word string) *model.Vocabulary {
	// 直接返回小于等于长度为 2 的单词
	if len(word) <= 2 {
		return &model.Vocabulary{Name: word}
	}

	// TODO 先获取单词原形，还原复数、比较级、过去式等单词形式

	// 提取所有疑似词根、前缀、后缀
	roots := extractSuspectedRoots(word)
	prefixes := extractSuspectedPrefixes(word)
	suffixes := extractSuspectedSuffixes(word)

	// 将所有词根、前缀、后缀进行组合
	results := make([]model.Vocabulary, 0)
	for _, root := range roots {
		for _, prefix := range prefixes {
			// 过滤掉组合长度超过目标单词长度的组合
			if len(prefix+root) > len(word) {
				continue
			}

			results = append(results, model.Vocabulary{Name: word, Roots: root, Prefix: prefix})
		}
	}

	for _, root := range roots {
		for _, suffix := range suffixes {
			// 过滤掉组合长度超过目标单词长度的组合
			if len(root+suffix) > len(word) {
				continue
			}

			results = append(results, model.Vocabulary{Name: word, Roots: root, Suffix: suffix})
		}
	}

	for _, root := range roots {
		for _, prefix := range prefixes {
			for _, suffix := range suffixes {
				// 过滤掉组合长度超过目标单词长度的组合
				if len(prefix+root+suffix) > len(word) || len(prefix+root) > len(word) || len(root+suffix) > len(word) {
					continue
				}

				results = append(results, model.Vocabulary{Name: word, Roots: root, Prefix: prefix, Suffix: suffix})
			}
		}
	}

	// 开始处理各种情况
	// 1. 理想情况：如果拼起来的正好是目标单词则直接返回
	greatResult, ok := lo.Find(results, func(item model.Vocabulary) bool {
		if item.Prefix+item.Roots+item.Suffix == item.Name {
			return true
		}

		if item.Roots+item.Suffix == item.Name {
			return true
		}

		if item.Prefix+item.Roots == item.Name {
			return true
		}

		return false
	})
	// stalling: error=stal  ling, good=stall ing
	// protestation error=prot est ation, good=pro test ation
	if ok {
		return &greatResult
	}

	// 2. 非理想情况：拼起来的长度小于目标单词长度
	sort.Slice(results, func(i, j int) bool {
		v1, v2 := results[i], results[j]
		return len(v1.Prefix+v1.Roots+v1.Suffix) > len(v2.Prefix+v2.Roots+v2.Suffix)
	})
	v, ok := lo.Find(results, func(item model.Vocabulary) bool { return len(word)-len(item.Prefix+item.Roots+item.Suffix) < 3 })
	if ok {
		return &v
	}

	// fmt.Println("====unknown====", word, results)

	// 没有匹配到任何
	return &model.Vocabulary{Name: word}
}

// extractSuspectedRoots
func extractSuspectedRoots(word string) []string {
	roots := lo.Filter(gRootAffixes, func(item string, index int) bool { return !strings.Contains(item, "-") })
	sort.Slice(roots, func(i, j int) bool { return roots[i] > roots[j] })
	return lo.Filter(roots, func(item string, index int) bool { return strings.Contains(word, item) })
}

// extractSuspectedPrefixes
func extractSuspectedPrefixes(word string) []string {
	prefixes := lo.Filter(gRootAffixes, func(item string, index int) bool { return strings.HasSuffix(item, "-") })
	suspected := lo.Filter(prefixes, func(item string, index int) bool { return strings.HasPrefix(word, strings.TrimSuffix(item, "-")) })
	sort.Slice(suspected, func(i, j int) bool { return suspected[i] < suspected[j] })
	return lo.Map(suspected, func(item string, index int) string { return strings.Trim(item, "-") })
}

// extractSuspectedSuffixes
func extractSuspectedSuffixes(word string) []string {
	roots := lo.Filter(gRootAffixes, func(item string, index int) bool { return strings.HasPrefix(item, "-") })
	suspected := lo.Filter(roots, func(item string, index int) bool { return strings.HasSuffix(word, strings.TrimPrefix(item, "-")) })
	sort.Slice(suspected, func(i, j int) bool { return suspected[i] < suspected[j] })
	return lo.Map(suspected, func(item string, index int) string { return strings.Trim(item, "-") })
}
