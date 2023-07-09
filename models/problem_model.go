package models

import (
	"github.com/imtiaz246/codera_oj/models/db"
	"gorm.io/gorm"
)

const (
	ProblemVisibilityStateProtected = iota
	ProblemVisibilityStatePublic
)

type Problem struct {
	gorm.Model
	UserID                           uint
	User                             User
	ContestID                        uint
	Contest                          Contest
	Title                            string
	TimeLimit                        float64
	MemoryLimit                      float64
	ProblemStatement                 string
	InputStatement                   string
	OutputStatement                  string
	NoteStatement                    string
	IsStatementsVisibleDuringContest bool
	ProblemTags                      []ProblemTag
	ProblemVisibilityState           uint `gorm:"default:0"`
}

func init() {
	if err := db.MigrateModelTables(Problem{}); err != nil {
		panic(err)
	}
}
