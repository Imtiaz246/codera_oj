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
	ContestID *uint    `json:"contestID"`
	Contest   *Contest `json:"contest"`
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
	Shares               []ProblemShare
	ChangeLogs           []ProblemChangeLog
	Discussions          []ProblemDiscussion
}

func init() {
	if err := db.MigrateModelTables(Problem{}); err != nil {
		panic(err)
	}
}

func (p *Problem) FindSharedUser(userID uint) *ProblemShare {
	for _, ps := range p.Shares {
		if ps.UserID == userID {
			return &ps
		}
	}
	return nil
}

func (p *Problem) GetSharedWithColumnName() string {
	return "Shares"
}

func (p *Problem) GetSharedUserColumnName() string {
	return "Shares.SharedWith"
}
