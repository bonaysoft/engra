package main

import (
	"github.com/bonaysoft/engra/pkg/dal/model"
	"gorm.io/gen"
)

func main() {
	g := gen.NewGenerator(gen.Config{
		OutPath: "./pkg/dal/query",
		Mode:    gen.WithoutContext | gen.WithDefaultQuery | gen.WithQueryInterface, // generate mode
	})

	// dsn := "root:admin@tcp(127.0.0.1:3306)/dicts?charset=utf8mb4&parseTime=True&loc=Local"
	// gormdb, _ := gorm.Open(mysql.Open(dsn))
	// g.UseDB(gormdb) // reuse your gorm db

	// Generate basic type-safe DAO API for struct `model.User` following conventions
	g.ApplyBasic(model.Vocabulary{}, model.RootsAffixes{}, model.EcDictWord{})

	// Generate the code
	g.Execute()
}
