package dict

import (
	"embed"
	_ "embed"
	"encoding/json"
	"fmt"
	"io/fs"
	"log"
	"os"
	"path/filepath"
	"strings"

	"github.com/bonaysoft/engra/pkg/dal/model"
	"github.com/bonaysoft/engra/pkg/dal/query"
	"github.com/gocarina/gocsv"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

//go:embed words.csv
var wordCsv []byte

func GetWords() []model.Vocabulary {
	words := make([]model.Vocabulary, 0)
	gocsv.UnmarshalBytes(wordCsv, &words)
	return words
}

//go:embed glossaries/*
var glossaryFs embed.FS

func Glossaries2Words() error {
	wm := NewWordsMgr()

	allWords := make([]model.Vocabulary, 0)
	fs.WalkDir(glossaryFs, ".", func(path string, d fs.DirEntry, err error) error {
		fmt.Println(path, d.Name())
		if d.IsDir() {
			return nil
		}
		bs, err := os.ReadFile(path)
		if err != nil {
			return err
		}

		var words []string
		if err := json.Unmarshal(bs, &words); err != nil {
			return err
		}

		for _, word := range words {
			wm.Add(word, strings.TrimSuffix(d.Name(), filepath.Ext(d.Name())))
		}

		allWords = append(allWords, wm.words...)
		return nil
	})

	f, err := os.Create("words.csv")
	if err != nil {
		return nil
	}

	return gocsv.MarshalFile(&allWords, f)
}

type WordsMgr struct {
	words []model.Vocabulary

	q  *query.Query
	dr *Roots
}

func NewWordsMgr() *WordsMgr {
	dsn := "/Users/yanbo/Develop/oss/ecdict/ecdict.db"
	gormdb, _ := gorm.Open(sqlite.Open(dsn), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Silent),
	})

	dr, err := NewRoots()
	if err != nil {
		log.Println(err)
	}

	return &WordsMgr{
		q:  query.Use(gormdb),
		dr: dr,
	}
}

func (wm *WordsMgr) Add(wordStr string, tags ...string) {
	for idx, word := range wm.words {
		if word.Name == wordStr {
			oldTags := strings.Split(word.Tags, ",")
			wm.words[idx].Tags = strings.Join(append(oldTags, tags...), ",")
			return
		}
	}
	w, err := wm.q.EcDictWord.Where(wm.q.EcDictWord.Word.Eq(wordStr)).Take()
	if err != nil {
		log.Println(fmt.Errorf("error find vocabulary %s: %v", wordStr, err))
	}

	v := model.Vocabulary{
		Name: wordStr,
		Tags: strings.Join(tags, ","),
	}

	_, wr, _ := wm.dr.Find(wordStr)
	if wr != nil {
		v.Roots = wr.Name
	}

	if w != nil {
		v.Phonetic = w.Phonetic
		v.Meaning = strings.ReplaceAll(w.Translation, "\n", "\\n")
		v.Exchange = strings.ReplaceAll(w.Exchange, "'", "\\'")
	}

	wm.words = append(wm.words, v)
}
