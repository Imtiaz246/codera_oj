package models

import (
	"github.com/imtiaz246/codera_oj/models/db"
	"gorm.io/gorm"
)

type CheckerType string

const (
	Default CheckerType = "Default"
	String  CheckerType = "String"
	Float   CheckerType = "Float"
	Special CheckerType = "Special"
)

type Problem struct {
	gorm.Model
	AuthorID  uint
	Author    *User    `gorm:"required;foreignKey:AuthorID" json:"-"`
	ContestID *uint    `json:"contest_id,omitempty"`
	Contest   *Contest `json:"-"`
	Title     string   `gorm:"Index;not null;type:varchar(55)"`

	TimeLimit                         float64
	MemoryLimit                       float64
	Statement                         string
	InputStatement                    string
	OutputStatement                   string
	NoteStatement                     string
	StatementsVisibilityDuringContest bool

	Tags      []ProblemTag
	Datasets  []Dataset
	Solutions []ProblemSolution

	IsProblemPublishable bool        `gorm:"default:false"`
	IsProblemPublished   bool        `gorm:"default:false"`
	CheckerType          CheckerType `gorm:"default:Default"`
	SharedWith           []ProblemShare
	ChangeLogs           []ProblemChangeLog
	Discussions          []ProblemDiscussion
}

func init() {
	if err := db.MigrateModelTables(Problem{}); err != nil {
		panic(err)
	}
}
