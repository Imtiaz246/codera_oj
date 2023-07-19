package v1

import "github.com/imtiaz246/codera_oj/models"

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
