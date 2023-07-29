package models

import (
	"github.com/imtiaz246/codera_oj/models/db"
	"gorm.io/gorm"
)

type ProblemSpecialChecker struct {
	gorm.Model
	Code       string
	LanguageID uint
	Language   *Language

	UserID        uint
	User          *User
	ProblemId     uint
	Problem       *Problem
	OwnerShipType PermitType
}

func init() {
	if err := db.MigrateModelTables(ProblemSpecialChecker{}); err != nil {
		panic(err)
	}
}
