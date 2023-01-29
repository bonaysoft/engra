package main

import (
	"encoding/json"
	"fmt"
	"io/fs"
	"path/filepath"
	"strings"

	"github.com/bitfield/script"
	"github.com/bonaysoft/engra/pkg/dal/model"
	"github.com/bonaysoft/engra/pkg/dal/query"
	"github.com/bonaysoft/engra/pkg/ra"
	"github.com/samber/lo"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func main() {
	dsn := "root:admin@tcp(127.0.0.1:3306)/dicts?charset=utf8mb4&parseTime=True&loc=Local"
	gormdb, _ := gorm.Open(mysql.Open(dsn), &gorm.Config{
		// Logger: logger.Default.LogMode(logger.Silent),
	})
	gormdb.AutoMigrate(model.Vocabulary{}, model.RootsAffixes{})
	q := query.Use(gormdb)

	// saveRa(q)

	// walkWords(q)

	// fillWordRoots(q)

	extractRa(q)

}

func saveRa(q *query.Query) {
	rs, err := q.RootsAffixes.Find()
	if err != nil {
		return
	}
	rootAffixes := lo.Map(rs, func(item *model.RootsAffixes, index int) string {
		if strings.HasPrefix(item.Name, "-") && strings.HasSuffix(item.Name, "-") {
			return strings.Trim(item.Name, "-")
		}

		return item.Name
	})
	script.Slice(rootAffixes).WriteFile("./pkg/ra/ra.txt")
}

func extractRa(q *query.Query) {
	vs, err := q.Vocabulary.Where(q.Vocabulary.Tag.Zero()).Find()
	if err != nil {
		return
	}

	var idx int
	for _, v := range vs {
		v := ra.Extract(v.Name)
		if v.Root == "" && v.Prefix == "" && v.Suffix == "" {
			continue
		}

		idx++
		fmt.Println(v.Name, v.Root, v.Prefix, v.Suffix)
	}
	fmt.Println(idx)
}

func fillWordRoots(q *query.Query) {
	contents, err := script.File("wordroots.txt").String()
	if err != nil {
		return
	}

	var idx, pIdx, sIdx int
	for _, v := range strings.Split(contents, "\n") {
		items := strings.Split(v, "--")
		vv, meaning := strings.TrimSpace(items[0]), strings.TrimSpace(items[1])
		if strings.Contains(meaning, "词根") {
			idx++
			vv = "-" + vv + "-"
			meaning = strings.Trim(meaning, "【词根】：")
		} else if (strings.Contains(meaning, "前缀") && strings.Contains(meaning, "后缀")) ||
			(strings.Contains(meaning, "前缀") && strings.Contains(meaning, "词根")) ||
			(strings.Contains(meaning, "后缀") && strings.Contains(meaning, "词根")) {
			fmt.Println(vv, meaning)
		} else if strings.Contains(meaning, "前缀") {
			pIdx++
			vv = vv + "-"
			meaning = strings.Trim(meaning, "【前缀】前缀：")
		} else if strings.Contains(meaning, "后缀") {
			sIdx++
			vv = "-" + vv
			meaning = strings.Trim(meaning, "【后缀】：")
		} else {
			fmt.Println(vv, meaning)
		}
		continue

		s := strings.Index(meaning, "/")
		e := strings.LastIndex(meaning, "/")
		if e > s {
			meaning = meaning[e+1:]
		}

		if strings.Contains(vv, "--") {
			fmt.Println(vv)
			return
		}

		meaning = strings.TrimSpace(meaning)
		meaning = strings.ReplaceAll(meaning, "【词根】", "；")
		meaning = strings.ReplaceAll(meaning, "词根：", "；")
		meaning = strings.ReplaceAll(meaning, "【前缀】", "；")
		meaning = strings.ReplaceAll(meaning, "【后缀】", "；")
		fmt.Println(vv, "==", meaning)

		v, err := q.RootsAffixes.Where(q.RootsAffixes.Name.Eq(vv)).Take()
		if err != nil {
			// return
		}

		v = &model.RootsAffixes{Name: vv, Meaning: meaning}
		if err := q.RootsAffixes.Where(q.RootsAffixes.Id.Eq(v.Id)).Save(v); err != nil {
		}
	}
	fmt.Println(idx, pIdx, sIdx)
}

func getRoot(root string, word string) string {
	roots := strings.Split(root, "=")
	if len(roots) == 1 {
		return root
	}

	for _, r := range roots {
		cr := strings.Trim(r, "-")
		if strings.Index(cr, "(") > 0 {
			cr = cr[:strings.Index(cr, "(")]
		}

		if strings.Contains(word, cr) {
			return r
		}
	}

	return ""
}

func walkWords2(q *query.Query) {
	contents, err := script.File("3.txt").String()
	if err != nil {
		return
	}

	var idx int
	var root, meaning, etymology, extension string
	rows := strings.Split(contents, "\n")
	for _, row := range rows {
		if strings.HasPrefix(row, "\t•\t") {
			idx++
			roots := strings.Split(strings.Trim(row, "\t•\t"), " ")
			root, meaning = roots[0], roots[1]
		} else if strings.HasPrefix(row, "【词源】") {
			etymology = strings.TrimPrefix(row, "【词源】")
		} else if strings.HasPrefix(row, "【引申】") {
			extension = strings.TrimPrefix(row, "【引申】")
			fmt.Sprintln(idx, root, meaning, etymology, extension)
		} else if strings.TrimSpace(row) != "" {
			word := strings.ToLower(strings.TrimSpace(strings.Split(row, "[")[0]))
			v, err := q.Vocabulary.Where(q.Vocabulary.Name.Eq(word)).Take()
			if err != nil {
				// return
				v = &model.Vocabulary{Name: word}
			}

			v.Root = getRoot(root, word)
			// fmt.Println(idx, v.Root, word)
			if v.Root == "" {
				fmt.Println("not found root", root, word)
			}

			if err := q.Vocabulary.Where(q.Vocabulary.Id.Eq(v.Id)).Save(v); err != nil {
				fmt.Println(err)
				return
			}
		}
	}

	fmt.Println(idx)
}

func walkWords(q *query.Query) {
	err := filepath.WalkDir("words", func(path string, d fs.DirEntry, err error) error {
		if d.IsDir() {
			return nil
		}

		tag := strings.Trim(path, "words/ .json")
		content, err := script.File(path).String()
		if err != nil {
			return err
		}

		w := make([]string, 0)
		if err := json.Unmarshal([]byte(content), &w); err != nil {
			return err
		}

		for _, s := range w {
			v, err := q.Vocabulary.Where(q.Vocabulary.Name.Eq(s)).Take()
			if err != nil {
				// return
				v = &model.Vocabulary{Name: s, Tag: tag}
			} else {
				v.Tag = v.Tag + "," + tag
			}

			if err := q.Vocabulary.Where(q.Vocabulary.Id.Eq(v.Id)).Save(v); err != nil {
				return err
			}
		}

		fmt.Println(path, len(w))
		return nil

	})
	fmt.Println(err)
}
