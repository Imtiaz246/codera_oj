package models

import (
	"github.com/imtiaz246/codera_oj/models/db"
	"gorm.io/gorm"
)

type Tag struct {
	gorm.Model
	TagName string `gorm:"unique;index"`
	UserID  uint
	AddedBy *User `gorm:"foreignKey:UserID"`
}

func init() {
	if err := db.MigrateModelTables(Tag{}); err != nil {
		panic(err)
	}
}
