package models

import (
	"github.com/imtiaz246/codera_oj/models/db"
	"gorm.io/gorm"
)

type Tag struct {
	gorm.Model
	name string
}

func init() {
	if err := db.MigrateModelTables(Tag{}); err != nil {
		panic(err)
	}
}
