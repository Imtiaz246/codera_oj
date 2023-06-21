package models

import (
	"github.com/imtiaz246/codera_oj/models/db"
	"gorm.io/gorm"
)

type Tags struct {
	gorm.Model
	name string
}

func init() {
	if err := db.MigrateModelTables(Tags{}); err != nil {
		panic(err)
	}
}
