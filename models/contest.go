package models

import (
	"github.com/imtiaz246/codera_oj/models/db"
	"gorm.io/gorm"
)

type Contest struct {
	gorm.Model
	// todo: complete
}

func init() {
	if err := db.MigrateModelTables(Contest{}); err != nil {
		panic(err)
	}
}
