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

func (t PermitType) IsPermitTypeValid() bool {
	return t == Editor || t == Viewer || t == Tester
}

type ProblemShare struct {
	gorm.Model
	UserID         uint
	SharedWith     *User `gorm:"foreignKey:UserID"`
	ProblemID      uint
	Problem        *Problem
	PermissionType PermitType `gorm:"default:Viewer"`
}

func (ps *ProblemShare) CanAddDataset() bool {
	return ps.PermissionType == Author || ps.PermissionType == Editor
}

func init() {
	if err := db.MigrateModelTables(ProblemShare{}); err != nil {
		panic(err)
	}
}
