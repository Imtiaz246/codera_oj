package models

import (
	"github.com/imtiaz246/codera_oj/models/db"
	"gorm.io/gorm"
)

type PermitType string

const (
	Author PermitType = "Author"
	Editor PermitType = "Editor"
	Viewer PermitType = "Viewer"
	Tester PermitType = "Tester"
)

type ProblemShare struct {
	gorm.Model
	UserID        uint
	User          *User
	ProblemID     uint
	Problem       *Problem
	OwnerShipType PermitType `gorm:"default:0"`
}

func init() {
	if err := db.MigrateModelTables(ProblemShare{}); err != nil {
		panic(err)
	}
}
