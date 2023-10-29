package problem

import (
	"gorm.io/gorm"
)

type ProblemTag struct {
	gorm.Model
	TagID     uint `json:"tagID"`
	Tag       *Tag
	ProblemID uint
	Problem   *Problem
}
