package dict

import (
	"fmt"
	"sort"
	"strconv"
)

type Item struct {
	Count     int
	RootCount int
}

type LevelCounter struct {
	m map[string]*Item
}

func NewLevelCounter() *LevelCounter {
	return &LevelCounter{m: make(map[string]*Item)}
}

func (c *LevelCounter) Count(tags []string, hitRoot bool) {
	for _, tag := range tags {
		v, ok := c.m[tag]
		if !ok {
			c.m[tag] = &Item{Count: 0}
		}

		if ok {
			v.Count++
		}
		if ok && hitRoot {
			v.RootCount++
		}
	}
}

func (c *LevelCounter) BuildSummary() [][]string {
	rows := make([][]string, 0)
	for name, item := range c.m {
		rows = append(rows, []string{name, strconv.Itoa(item.Count),
			fmt.Sprintf("%.2f%%", float64(item.RootCount)/float64(item.Count)*100)})
	}

	sort.Slice(rows, func(i, j int) bool { return rows[i][0] < rows[j][0] })
	return rows
}
