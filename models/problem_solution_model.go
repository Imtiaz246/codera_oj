package models

import (
	"github.com/imtiaz246/codera_oj/models/db"
	"gorm.io/gorm"
	"time"
)

type ProblemSolution struct {
	gorm.Model
	Code         string
	LanguageID   uint
	Language     *Language
	LastExecuted time.Time
	TimeTaken    float32
	MemoryTaken  float64

	UserID        uint
	User          *User
	ProblemId     uint
	Problem       *Problem
	OwnerShipType PermitType
}

func init() {
	if err := db.MigrateModelTables(ProblemSolution{}); err != nil {
		panic(err)
	}
}
