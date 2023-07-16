package models

import (
	"github.com/imtiaz246/codera_oj/models/db"
	"gorm.io/gorm"
)

const (
	CheckerTypeDefault = iota
	CheckerTypeString
	CheckerTypeFloat
	CheckerTypeSpecial
)

type Problem struct {
	gorm.Model
	UserID    uint
	User      *User
	ContestID uint
	Contest   *Contest
	Title     string

	TimeLimit                         float64
	MemoryLimit                       float64
	Statement                         string
	InputStatement                    string
	OutputStatement                   string
	NoteStatement                     string
	StatementsVisibilityDuringContest bool

	Tags     []ProblemTag
	Datasets []Dataset

	CheckerType          uint `gorm:"default:0"`
	IsProblemPublishable bool `gorm:"default:false"`
	SharedWith           []ProblemShare
	ChangeLogs           []ProblemChangeLog
	Discussions          []ProblemDiscussion
}

func init() {
	if err := db.MigrateModelTables(Problem{}); err != nil {
		panic(err)
	}
}
