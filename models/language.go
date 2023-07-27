package models

import (
	"github.com/imtiaz246/codera_oj/models/db"
	"gorm.io/gorm"
)

type Language struct {
	gorm.Model
	Name string `gorm:"unique;required"`
}

func init() {
	if err := db.MigrateModelTables(Language{}); err != nil {
		panic(err)
	}
}
