package models

import (
	"github.com/imtiaz246/codera_oj/models/db"
	"gorm.io/gorm"
)

type OwnershipType int

const (
	Author OwnershipType = iota
	Editor
	Viewer
	Tester
)

func (i OwnershipType) String() string {
	return [...]string{"Author", "Editor", "Viewer", "Tester"}[i-1]
}

type ProblemShare struct {
	gorm.Model
	UserID        uint
	User          *User
	ProblemID     uint
	Problem       *Problem
	OwnerShipType OwnershipType `gorm:"default:0"`
}

func init() {
	if err := db.MigrateModelTables(ProblemShare{}); err != nil {
		panic(err)
	}
}
