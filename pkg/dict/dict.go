package dict

import (
	"bytes"
	"fmt"
	"os"
	"regexp"
	"strings"

	"github.com/bitfield/script"
	"github.com/bonaysoft/engra/dict"
	"github.com/olekukonko/tablewriter"
)

type Dict struct {
	roots        *Roots
	LevelCounter *LevelCounter
}

func NewDict() (*Dict, error) {
	roots, err := NewRoots()
	if err != nil {
		return nil, err
	}

	return &Dict{
		roots:        roots,
		LevelCounter: NewLevelCounter(),
	}, nil
}

// BuildWordsMd 将词库单词与词根进行匹配，然后将词库数据汇总到一个md文件中
func (d *Dict) BuildWordsMd() error {
	var count float64
	rows := make([][]string, 0)
	words := dict.GetWords()
	for idx, word := range words {
		wr, _ := d.roots.Find(word.Name)
		if wr != nil {
			if words[idx].Root != "" {
				words[idx].Root += "," + wr.Name
			} else {
				words[idx].Root += wr.Name
			}
			v, _ := wr.Find(word.Name)
			words[idx].Status = v.Status()
			count++
		}

		rows = append(rows, words[idx].Row())
	}

	f, err := os.Create("dict/words.md")
	if err != nil {
		return err
	}

	fmt.Fprintf(f, "# 汇总\n\n")
	fmt.Fprintf(f, "- 总单词量：%d\n", len(words))
	fmt.Fprintf(f, "- 已关联词根单词数：%.f\n", count)
	fmt.Fprintf(f, "- 词根覆盖度：%.2f%%\n\n", count/float64(len(words))*100)
	fmt.Fprintln(f, "## 词汇表")
	table := tablewriter.NewWriter(f)
	table.SetHeader([]string{"name", "tags", "roots", "status"})
	table.SetBorders(tablewriter.Border{Left: true, Top: false, Right: true, Bottom: false})
	table.SetCenterSeparator("|")
	table.AppendBulk(rows) // Add Bulk Data
	table.Render()
	return nil
}

func (d *Dict) BuildSummary() error {
	words := dict.GetWords()
	for _, word := range words {
		tags := strings.Split(word.Tag, ",")
		_, err := d.roots.Find(word.Name)
		d.LevelCounter.Count(tags, err == nil)
	}

	buf := bytes.NewBuffer([]byte{})
	table := tablewriter.NewWriter(buf)
	table.SetHeader([]string{"名称", "单词量", "词根覆盖度"})
	table.SetBorders(tablewriter.Border{Left: true, Top: false, Right: true, Bottom: false})
	table.SetAlignment(tablewriter.ALIGN_LEFT)
	table.SetCenterSeparator("|")
	table.AppendBulk(d.LevelCounter.BuildSummary()) // Add Bulk Data
	table.Render()
	return replaceReadmeSection(buf.String())
}

func replaceReadmeSection(txt string) error {
	txt = "<!--START_SECTION:engra-->\n\n" + txt + "\n<!--END_SECTION:engra-->"
	s, err := script.File("README.md").String()
	if err != nil {
		return nil
	}

	exp := regexp.MustCompile(`<!--START_SECTION.+([\s\S]+\n).+END_SECTION:engra-->`)
	_, err = script.Echo(exp.ReplaceAllString(s, txt)).WriteFile("README.md")
	return err
}
