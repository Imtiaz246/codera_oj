package v1

import (
	"github.com/imtiaz246/codera_oj/models"
	"gorm.io/gorm"
)

type CreateProblemOption struct {
	Title string `json:"title" validate:"required"`
}

type UpdateProblemOption struct {
	TimeLimit                         float64            `json:"timeLimit"`
	MemoryLimit                       float64            `json:"memoryLimit"`
	Statement                         string             `json:"statement"`
	InputStatement                    string             `json:"inputStatement"`
	OutputStatement                   string             `json:"outputStatement"`
	NoteStatement                     string             `json:"noteStatement"`
	StatementsVisibilityDuringContest bool               `json:"statementsVisibilityDuringContest"`
	CheckerType                       models.CheckerType `json:"checkerType"`
}

func (p *UpdateProblemOption) UpdateProblemModelFormat(user *models.User, ID uint) *models.Problem {
	return &models.Problem{
		Model:                             gorm.Model{ID: ID},
		Author:                            user,
		TimeLimit:                         p.TimeLimit,
		MemoryLimit:                       p.MemoryLimit,
		Statement:                         p.Statement,
		InputStatement:                    p.InputStatement,
		OutputStatement:                   p.OutputStatement,
		NoteStatement:                     p.NoteStatement,
		StatementsVisibilityDuringContest: p.StatementsVisibilityDuringContest,
		CheckerType:                       p.CheckerType,
	}
}

type AddTagOption struct {
	TagName string `json:"tagName"`
}
