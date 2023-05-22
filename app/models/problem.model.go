package models

import (
	"gorm.io/gorm"
)

const (
	ProblemVisibilityStatePrivate = iota
	ProblemVisibilityStatePublic
)

type Problem struct {
	gorm.Model
	UserID                            uint
	User                              User
	ContestID                         uint
	Contest                           Contest
	Title                             string
	TimeLimit                         float64
	MemoryLimit                       float64
	ProblemStatement                  string
	InputStatement                    string
	OutputStatement                   string
	NoteStatement                     string
	VisibilityStatementsDuringContest bool
	TagsID                            uint
	Tags                              Tags
	ProblemVisibilityState            uint `gorm:"default:0"`
}

type Tags struct {
	gorm.Model
	name string
}
