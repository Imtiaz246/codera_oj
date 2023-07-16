package models

import (
	"github.com/imtiaz246/codera_oj/models/db"
	"gorm.io/gorm"
)

type CheckerType int

const (
	Default CheckerType = iota
	String
	Float
	Special
)

func (i CheckerType) String() string {
	return [...]string{"Default", "String", "Float", "Special"}[i-1]
}

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

	IsProblemPublishable bool        `gorm:"default:false"`
	IsProblemPublished   bool        `gorm:"default:false"`
	CheckerType          CheckerType `gorm:"default:0"`
	SharedWith           []ProblemShare
	ChangeLogs           []ProblemChangeLog
	Discussions          []ProblemDiscussion
}

func init() {
	if err := db.MigrateModelTables(Problem{}); err != nil {
		panic(err)
	}
}
