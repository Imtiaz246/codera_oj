package problem

import (
	"github.com/imtiaz246/codera_oj/internal/core/domain/models"
	"gorm.io/gorm"
)

const (
	Default CheckerType = "Default"
	String  CheckerType = "String"
	Float   CheckerType = "Float"
	Special CheckerType = "Special"
)

type CheckerType string

type Problem struct {
	gorm.Model
	AuthorID  uint
	Author    *models.User    `gorm:"required;foreignKey:AuthorID" json:"-"`
	ContestID *uint           `json:"contestID"`
	Contest   *models.Contest `json:"contest"`
	// todo: add unique with is problem published
	Title string `gorm:"Index;not null;type:varchar(55)"`

	TimeLimit                         float64
	MemoryLimit                       float64
	Statement                         string
	InputStatement                    string
	OutputStatement                   string
	NoteStatement                     string
	StatementsVisibilityDuringContest bool

	Tags      []ProblemTag
	Datasets  []Dataset
	Solutions []Solution

	IsProblemPublishable bool        `gorm:"default:false"`
	IsProblemPublished   bool        `gorm:"default:false"`
	CheckerType          CheckerType `gorm:"default:Default"`
	SharedWith           []Share
	ChangeLogs           []ChangeLog
	Discussions          []Discussion
}
