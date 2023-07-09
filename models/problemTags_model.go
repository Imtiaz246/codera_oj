package models

import (
	"github.com/imtiaz246/codera_oj/models/db"
	"gorm.io/gorm"
)

type ProblemTag struct {
	gorm.Model
	TagID     uint
	Tag       Tag
	ProblemID uint
	Problem   Problem
}

func init() {
	if err := db.MigrateModelTables(ProblemTag{}); err != nil {
		panic(err)
	}
}
